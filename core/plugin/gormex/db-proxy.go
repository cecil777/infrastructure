package gormex

import (
	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
)

type dbProxy struct {
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
		s.db = d
	}

	return s.db, nil
}
