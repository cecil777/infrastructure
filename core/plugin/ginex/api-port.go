package ginex

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/cecil777/infrastructure/core/api"
	"github.com/cecil777/infrastructure/core/db"
	"github.com/cecil777/infrastructure/core/dp/ioc"
	"github.com/cecil777/infrastructure/core/errorex"
	"github.com/cecil777/infrastructure/core/plugin/gormex"
	"github.com/cecil777/infrastructure/core/runtimeex"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"gorm.io/driver/mysql"
)

type routeParam struct {
	API      string `uri:"api"`
	Endpoint string `uri:"endpoint"`
}

type apiPort struct {
	apiFactory api.IFactory
	port       int
	req        *http.Request
	resp       http.ResponseWriter
}

func (m *apiPort) Listen() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	validate := validator.New()
	app.POST("/:endpoint/:api", func(ctx *gin.Context) {
		var rp routeParam
		ctx.ShouldBindUri(&rp)

		var resp api.Response
		resp.Data = ""

		var err error
		defer func() {
			if rv := recover(); rv != nil {
				if cErr, ok := rv.(error); ok {
					err = cErr
				} else {
					err = fmt.Errorf("%v", rv)
				}
			}

			if err != nil {
				if cErr, ok := err.(errorex.Custom); ok {
					resp.Error = int(cErr.Code)
					resp.Data = cErr.Error()
				} else {
					fmt.Printf("%v\n", err)
					resp.Error = int(errorex.PanicCode)
				}
			}

			ctx.JSON(http.StatusOK, resp)
		}()

		apiInstance := m.apiFactory.Build(rp.Endpoint, rp.API)
		ctx.BindJSON(apiInstance)
		tracer := bootTracerBasedJaeger()
		var parentSpan opentracing.Span

		spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		if err != nil {
			parentSpan = tracer.StartSpan(ctx.Request.URL.Path)
			defer parentSpan.Finish()
		} else {
			parentSpan = opentracing.StartSpan(
				ctx.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer parentSpan.Finish()
		}
		ctx.Set("Tracer", tracer)
		ctx.Set("ParentSpanContext", parentSpan)
		drive := mysql.Open("root:root1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
		ioc.Set(
			(*db.IFactory)(nil),
			gormex.NewFactory(ctx, drive),
		)
		ioc.Inject(apiInstance, func(v reflect.Value) reflect.Value {
			if traceable, ok := v.Interface().(runtimeex.ITraceable); ok {
				traceable.SetTraceID("")     // traceID注入
				traceable.SetTraceSpanID("") // traceSpanID注入
			}

			return v
		})
		if err = validate.Struct(apiInstance); err != nil {
			err = errorex.New(errorex.VerifyCode, "")
			return
		}

		resp.Data, err = apiInstance.Call()
	})

	if m.req != nil && m.resp != nil {
		app.ServeHTTP(m.resp, m.req)
	} else {
		fmt.Printf(
			"listen:%d[%s]\n",
			m.port,
			time.Now().Format("2006-01-02 15:04:05"),
		)
		app.Run(
			fmt.Sprintf(":%d", m.port),
		)
	}
}

// NewAPIPort is 创建gin端口实例
func NewAPIPort(apiFactory api.IFactory, port int) api.IPort {
	return &apiPort{
		apiFactory: apiFactory,
		port:       port,
	}
}

func bootTracerBasedJaeger() opentracing.Tracer {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		ServiceName: "http-server-test",
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			//LocalAgentHostPort:  "127.0.0.1:6381",
			BufferFlushInterval: 100 * time.Millisecond,
			CollectorEndpoint:   "http://127.0.0.1:14268/api/traces",
		},
	}

	tracer, _, err := cfg.NewTracer(
		config.Logger(jaegerLog.StdLogger),
		config.ZipkinSharedRPCSpan(true),
	)
	if err != nil {
		log.Printf("failed to use jaeger tracer plugin, got error %v", err)
		os.Exit(1)
	}

	return tracer
}
