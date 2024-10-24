package gormex

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
)

type dbProxy struct {
	dns string
	db  *gorm.DB
}

func (s *dbProxy) getDb() (*gorm.DB, error) {
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
