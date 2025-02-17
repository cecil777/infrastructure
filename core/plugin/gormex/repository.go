package gormex

import (
	"gorm.io/gorm"

	"github.com/cecil777/infrastructure/core/db"
)

type repository struct {
	model interface{}
	isTx  bool
	proxy dbProxy
	uow   *uowRepository
}

func (r repository) Query() db.IQuery {
	q := &query{
		proxy: r.proxy,
	}
	q.actions = append(q.actions, func(db *gorm.DB) *gorm.DB {
		return db.Model(r.model)
	})
	return q
}

func (r repository) Add(entry db.IIdentity) error {
	if r.isTx {
		r.uow.addQueue = append(r.uow.addQueue, entry)
		return nil
	}
	connDb, err := r.proxy.getDb()
	if err != nil {
		return err
	}
	err = connDb.Create(entry).Error
	return err
}

func (r repository) Remove(entry db.IIdentity) error {
	if r.isTx {
		r.uow.removeQueue = append(r.uow.removeQueue, entry)
		return nil
	}
	connDb, err := r.proxy.getDb()
	if err != nil {
		return err
	}
	err = connDb.Delete(entry).Error
	return err
}

func (r repository) Save(entry db.IIdentity) error {
	if r.isTx {
		r.uow.saveQueue = append(r.uow.saveQueue, entry)
		return nil
	}
	connDb, err := r.proxy.getDb()
	if err != nil {
		return err
	}
	err = connDb.Save(entry).Error
	return err
}
