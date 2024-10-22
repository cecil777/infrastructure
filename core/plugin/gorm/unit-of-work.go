package gorm

import (
	"gorm.io/gorm"
)

type uowRepository struct {
	tx     *gorm.DB
	add    []interface{}
	save   []interface{}
	remove []interface{}
}

func (u uowRepository) Commit() error {
	if len(u.add) > 0 || len(u.remove) > 0 || len(u.save) > 0 {
		tx := u.tx.Begin()
		if len(u.add) > 0 {
			for _, q := range u.add {
				if err := tx.Create(q).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		}
		if len(u.remove) > 0 {
			for _, q := range u.remove {
				if err := tx.Delete(q).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		}
		if len(u.save) > 0 {
			for _, q := range u.save {
				if err := tx.Save(q).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		}
		err := tx.Commit().Error
		if err != nil {
			return err
		}
	}
	return nil
}
