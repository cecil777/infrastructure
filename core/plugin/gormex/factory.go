package gormex

import (
	"github.com/cecil777/infrastructure/core/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
	"log"
)

type gormExFactory struct {
	dns string
	db  *gorm.DB
}

func NewFactory(dns string) db.IFactory {
	return &gormExFactory{dns: dns}
}

func (s *gormExFactory) getDb() (*gorm.DB, error) {
	if s.db == nil {
		var d *gorm.DB
		var err error
		if s.dns != "" {
			d, err = gorm.Open(mysql.Open(s.dns), &gorm.Config{})
		} else {
			d, err = gorm.Open(tests.DummyDialector{}, nil)
		}

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		s.db = d
	}

	return s.db, nil
}

func (s *gormExFactory) Db(entry db.IIdentity, extra ...interface{}) db.IRepository {
	r := &repository{}
	s.db, _ = s.getDb()
	r.db = s.db
	r.model = entry
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
	d, _ := s.getDb()
	return &uowRepository{
		tx: d,
	}
}
