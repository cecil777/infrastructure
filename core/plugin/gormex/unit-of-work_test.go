package gormex

import (
	"core/db"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnitOfWork(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockIUnitOfWork := db.NewMockIUnitOfWork(ctl)
	gomock.InOrder(mockIUnitOfWork.EXPECT().Commit().Return(nil))
	err := mockIUnitOfWork.Commit()
	assert.Equal(t, err, nil)
}
