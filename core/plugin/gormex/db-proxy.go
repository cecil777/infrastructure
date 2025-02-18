package gormex

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
	gormopentracing "gorm.io/plugin/opentracing"
)

type dbProxy struct {
	ctx   context.Context
	db    *gorm.DB
	drive gorm.Dialector
}

func (s *dbProxy) getDb() (*gorm.DB, error) {
	if s.db == nil {
		var d *gorm.DB
		var err error
		if s.drive != nil {
			d, err = gorm.Open(s.drive, &gorm.Config{})
		} else {
			d, err = gorm.Open(tests.DummyDialector{}, nil)
		}
		if err != nil {
			return nil, err
		}
		if s.ctx != nil {
			parentContext := s.ctx.Value("ParentSpanContext")
			if sp, ok := parentContext.(opentracing.Span); ok {
				newCtx := opentracing.ContextWithSpan(s.ctx, sp)
				d = d.WithContext(newCtx)
				tracer := s.ctx.Value("Tracer")
				if engine, ok := tracer.(opentracing.Tracer); ok {
					opentracing.SetGlobalTracer(engine)
					err = d.Use(gormopentracing.New())
					if err != nil {
						return nil, err
					}
				}
			}
		}
		s.db = d
	}

	return s.db, nil
}
