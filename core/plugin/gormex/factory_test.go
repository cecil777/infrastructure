package gormex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactoryDb(t *testing.T) {
	test := Test{}
	g := NewFactory(nil, nil)
	gd := g.Db(test)
	repo, ok := gd.(*repository)
	assert.True(t, ok)
	assert.False(t, repo.isTx)
	assert.Equal(t, repo.model, test)
}

func TestFactoryUow(t *testing.T) {
	test := Test{}
	g := NewFactory(nil, nil)
	uow := g.Uow()
	_ = g.Db(test, uow)
	u, ok := uow.(*uowRepository)
	assert.True(t, ok)
	assert.Equal(t, len(u.addQueue), 0)
	assert.Equal(t, len(u.saveQueue), 0)
	assert.Equal(t, len(u.removeQueue), 0)
}
