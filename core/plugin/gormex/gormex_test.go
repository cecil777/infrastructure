package gormex

import (
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
