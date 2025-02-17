package db

type IQuery interface {
	Count() (int64, error)
	Order(fields ...string) IQuery
	OrderByDesc(fields ...string) IQuery
	Skip(v int) IQuery
	Take(v int) IQuery
	ToArray(dst interface{}) error
	Where(args ...interface{}) IQuery
}
