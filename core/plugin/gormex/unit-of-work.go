package gormex

type uowRepository struct {
	proxy       dbProxy
	addQueue    []interface{}
	saveQueue   []interface{}
	removeQueue []interface{}
}

func (u uowRepository) Commit() (err error) {
	if len(u.addQueue) > 0 || len(u.removeQueue) > 0 || len(u.saveQueue) > 0 {
		db, err := u.proxy.getDb()
		if err != nil {
			return err
		}

		tx := db.Begin()
		defer func() {
			if err != nil {
				tx.Rollback()
			}
		}()
		if len(u.addQueue) > 0 {
			for _, q := range u.addQueue {
				if err = tx.Create(q).Error; err != nil {
					return err
				}
			}
		}
		if len(u.removeQueue) > 0 {
			for _, q := range u.removeQueue {
				if err = tx.Delete(q).Error; err != nil {
					return err
				}
			}
		}
		if len(u.saveQueue) > 0 {
			for _, q := range u.saveQueue {
				if err = tx.Save(q).Error; err != nil {
					return err
				}
			}
		}
		err = tx.Commit().Error
	}
	return nil
}
