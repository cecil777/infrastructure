package db

// IFactory is 数据工厂
type IFactory interface {
	Db(entry IIdentity, extra ...interface{}) IRepository
	Uow() IUnitOfWork
}
