package gormex

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactoryDb(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	test := Test{}
	g := NewFactory("")
	gd := g.Db(test)
	repo, ok := gd.(*repository)
	assert.True(t, ok)
	assert.False(t, repo.isTx)
	assert.Equal(t, repo.model, test)
	assert.NotEqual(t, repo.db, nil)
}

func TestFactoryUow(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	test := Test{}
	g := NewFactory("")
	uow := g.Uow()
	_ = g.Db(test, uow)
	u, ok := uow.(*uowRepository)
	assert.True(t, ok)
	assert.Equal(t, len(u.add), 0)
	assert.Equal(t, len(u.save), 0)
	assert.Equal(t, len(u.remove), 0)
}
