package db

// IRepository is 表仓库
type IRepository interface {
	Add(entry IIdentity) error
	Query() IQuery
	Remove(entry IIdentity) error
	Save(entry IIdentity) error
}
