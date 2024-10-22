package gormex

import (
	"core/db"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestFactoryDb(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	//test := Test{}
	//方案1
	//g := NewFactory("")
	//gd := g.Db(test)
	//repo, ok := gd.(*repository)
	//assert.True(t, ok)
	//assert.False(t, repo.isTx)
	//assert.Equal(t, repo.model, test)
	//assert.NotEqual(t, repo.db, nil)

	//方案2
	factory := db.NewMockIFactory(ctl)
	identity := db.NewMockIIdentity(ctl)
	factory.EXPECT().Db(identity).Return(&repository{
		model: identity,
	})
	factory.Db(identity)
}

func TestFactoryUow(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	//test := Test{}
	//方案1
	//g := NewFactory("")
	//uow := g.Uow()
	//_ = g.Db(test, uow)
	//u, ok := uow.(*uowRepository)
	//assert.True(t, ok)
	//assert.Equal(t, len(u.add), 0)
	//assert.Equal(t, len(u.save), 0)
	//assert.Equal(t, len(u.remove), 0)
	//方案2
	factory := db.NewMockIFactory(ctl)
	factory.EXPECT().Uow().Return(&uowRepository{})
	factory.Uow()
}
