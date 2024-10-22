package gormex

import (
	"core/db"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryQuery(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	//test := Test{}
	//g := NewFactory("")
	//q, ok := g.Db(test).Query().(*query)
	//assert.True(t, ok)
	//assert.Equal(t, q.model, test)
	//assert.NotEqual(t, q.db, nil)

	mockIRepository := db.NewMockIRepository(ctl)
	mockIRepository.EXPECT().Query().Return(&query{})
	mockIRepository.Query()
}

func TestRepositoryAdd(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockIRepository := db.NewMockIRepository(ctl)
	mockIIdentity := db.NewMockIIdentity(ctl)
	gomock.InOrder(mockIRepository.EXPECT().Add(mockIIdentity).Return(nil))
	err := mockIRepository.Add(mockIIdentity)
	assert.Equal(t, err, nil)
}

func TestRepositorySave(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockIRepository := db.NewMockIRepository(ctl)
	mockIIdentity := db.NewMockIIdentity(ctl)
	gomock.InOrder(mockIRepository.EXPECT().Save(mockIIdentity).Return(nil))
	err := mockIRepository.Save(mockIIdentity)
	assert.Equal(t, err, nil)
}

func TestRepositoryRemove(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockIRepository := db.NewMockIRepository(ctl)
	mockIIdentity := db.NewMockIIdentity(ctl)
	gomock.InOrder(mockIRepository.EXPECT().Remove(mockIIdentity).Return(nil))
	err := mockIRepository.Remove(mockIIdentity)
	assert.Equal(t, err, nil)
}
