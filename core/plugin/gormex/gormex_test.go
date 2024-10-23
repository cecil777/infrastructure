package gormex

import (
	"core/db"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Test struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t Test) GetID() string {
	return strconv.Itoa(int(t.ID))
}

func (Test) TableName() string {
	return "test"
}

func NewMock() db.IFactory {
	return NewFactory("root:123456@tcp(127.0.0.1:3306)/yddb?charset=utf8mb4&parseTime=True&loc=Local")
}

func DeleteMockTest(conn db.IFactory) {
	at := Test{}
	factory := conn.(*gormExFactory)
	factory.db.Where("1 = 1").Delete(&at)
}

func MultipleCreateTest(conn db.IFactory, name string) {
	for i := 1; i < 4; i++ {
		at := Test{}
		at.Name = fmt.Sprintf("%s %d", name, i)
		c := conn.Db(at)
		_ = c.Add(&at)
	}
}
