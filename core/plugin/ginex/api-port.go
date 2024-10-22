package ginex

import (
	"core/api"
	"core/errorex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type routeParam struct {
	API      string `uri:"api"`
	Endpoint string `uri:"endpoint"`
}

type apiPort struct {
	port int
	req  *http.Request
	resp http.ResponseWriter
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

		apiInstance := api.Build(rp.Endpoint, rp.API)
		ctx.BindJSON(apiInstance)
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
func NewAPIPort(port int) api.IPort {
	return &apiPort{
		port: port,
	}
}
