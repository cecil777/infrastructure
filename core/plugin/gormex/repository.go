package gormex

import (
	"core/db"
	"gorm.io/gorm"
)

type repository struct {
	model interface{}
	isTx  bool
	db    *gorm.DB
	uow   *uowRepository
}

func (r repository) Query() db.IQuery {
	return &query{
		db:    r.db,
		model: r.model,
	}
}

func (r repository) Add(entry db.IIdentity) error {
	if r.isTx {
		r.uow.add = append(r.uow.add, entry)
		return nil
	}
	err := r.db.Create(entry).Error
	return err
}

func (r repository) Remove(entry db.IIdentity) error {
	if r.isTx {
		r.uow.remove = append(r.uow.remove, entry)
		return nil
	}
	err := r.db.Delete(entry).Error
	return err
}

func (r repository) Save(entry db.IIdentity) error {
	if r.isTx {
		r.uow.save = append(r.uow.save, entry)
		return nil
	}
	err := r.db.Save(entry).Error
	return err
}