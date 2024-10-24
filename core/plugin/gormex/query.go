package gormex

import (
	"fmt"

	"github.com/cecil777/infrastructure/core/db"

	"gorm.io/gorm"
)

type query struct {
	proxy   dbProxy
	actions []func(db *gorm.DB) *gorm.DB
}

func (q *query) Count() (int64, error) {
	var count int64
	connDb, err := q.proxy.getDb()
	if err != nil {
		return count, err
	}
	if len(q.actions) > 0 {
		for _, fn := range q.actions {
			connDb = fn(connDb)
		}
	}
	connDb.Count(&count)
	return count, err
}

func (q *query) Order(fields ...string) db.IQuery {
	if len(fields) > 0 {
		for _, field := range fields {
			q.actions = append(q.actions, func(db *gorm.DB) *gorm.DB {
				return db.Order(field)
			})
		}
	}
	return q
}

func (q *query) OrderByDesc(fields ...string) db.IQuery {
	if len(fields) > 0 {
		for _, field := range fields {
			q.actions = append(q.actions, func(db *gorm.DB) *gorm.DB {
				return db.Order(fmt.Sprintf("%s desc", field))
			})
		}
	}
	return q
}

func (q *query) Skip(v int) db.IQuery {
	q.actions = append(q.actions, func(db *gorm.DB) *gorm.DB {
		return db.Offset(v)
	})
	return q
}

func (q *query) Take(v int) db.IQuery {
	q.actions = append(q.actions, func(db *gorm.DB) *gorm.DB {
		return db.Limit(v)
	})
	return q
}

func (q *query) ToArray(dst interface{}) error {
	connDb, err := q.proxy.getDb()
	if err != nil {
		return err
	}
	if len(q.actions) > 0 {
		for _, fn := range q.actions {
			connDb = fn(connDb)
		}
	}
	err = connDb.Find(dst).Error
	return err
}

func (q *query) Where(args ...interface{}) db.IQuery {
	q.actions = append(q.actions, func(db *gorm.DB) *gorm.DB {
		return db.Where(args)
	})
	return q
}
