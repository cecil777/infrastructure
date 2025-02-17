package gormex

import (
	"github.com/cecil777/infrastructure/core/db"
	"gorm.io/gorm"
)

type gormExFactory struct {
	proxy dbProxy
}

func (s *gormExFactory) Db(entry db.IIdentity, extra ...interface{}) db.IRepository {
	r := &repository{}
	r.model = entry
	r.proxy = s.proxy
	if len(extra) > 0 {
		r.isTx = true
		if extra[0] != nil {
			uow := extra[0].(*uowRepository)
			r.uow = uow
		}
	}
	return r
}

func (s *gormExFactory) Uow() db.IUnitOfWork {
	return &uowRepository{
		proxy: s.proxy,
	}
}

func NewFactory(drive gorm.Dialector) db.IFactory {
	return &gormExFactory{proxy: dbProxy{drive: drive}}
}
