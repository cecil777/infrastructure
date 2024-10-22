package gormex

import (
	"core/db"
	"fmt"
	"gorm.io/gorm"
)

type query struct {
	db    *gorm.DB
	model interface{}
}

func (q query) Count() (int64, error) {
	var count int64
	q.db.Model(q.model).Where(q.model).Count(&count)
	return count, nil
}

func (q query) Order(fields ...string) db.IQuery {
	if len(fields) > 0 {
		for _, field := range fields {
			q.db = q.db.Order(field)
		}
	}
	return q
}

func (q query) OrderByDesc(fields ...string) db.IQuery {
	if len(fields) > 0 {
		for _, field := range fields {
			q.db = q.db.Order(fmt.Sprintf("%s desc", field))
		}
	}
	return q
}

func (q query) Skip(v int) db.IQuery {
	q.db = q.db.Offset(v)
	return q
}

func (q query) Take(v int) db.IQuery {
	q.db = q.db.Limit(v)
	return q
}

func (q query) ToArray(dst interface{}) error {
	q.db = q.db.Find(dst)
	return nil
}

func (q query) Where(args ...interface{}) db.IQuery {
	q.db = q.db.Where(args)
	return q
}
