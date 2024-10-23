package gormex

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryQuery(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	test := Test{}
	g := NewFactory("")
	_, ok := g.Db(test).Query().(*query)
	assert.True(t, ok)
}

func TestRepositoryAdd(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	first := Test{}
	c := conn.Db(first)
	first.Name = "repository add test"
	err := c.Add(&first)
	assert.NoError(t, err)
	assert.NotEqual(t, first.ID, uint(0))
	assert.Equal(t, first.Name, "repository add test")
	DeleteMockTest(conn)
}

func TestRepositorySave(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	first := Test{}
	c := conn.Db(first)
	first.Name = "repository save test"
	err := c.Add(&first)
	first.Name = "repository update test"
	err = c.Save(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.Name, "repository update test")
	DeleteMockTest(conn)
}

func TestRepositoryRemove(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	first := Test{}
	c := conn.Db(first)
	first.Name = "repository remove test"
	err := c.Add(&first)
	assert.NotEqual(t, first.ID, uint(0))
	err = c.Remove(&first)
	assert.NoError(t, err)
	assert.True(t, first.DeletedAt.Valid)
	DeleteMockTest(conn)
}

func TestUowRepositoryAdd(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	first := Test{}
	uow := conn.Uow()
	c := conn.Db(first, uow)
	first.Name = "repository uow add test"
	err := c.Add(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.ID, uint(0))
	first.Name = "repository uow update test"
	err = c.Save(&first)
	err = uow.Commit()
	assert.NoError(t, err)
	assert.NotEqual(t, first.ID, uint(0))
	assert.Equal(t, first.Name, "repository uow update test")
	DeleteMockTest(conn)
}

func TestUowRepositorySave(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	first := Test{}
	c := conn.Db(first)
	first.Name = "repository uow save test"
	err := c.Add(&first)
	updatedAt := first.UpdatedAt
	uow := conn.Uow()
	cc := conn.Db(first, uow)
	first.Name = "repository uow update test"
	err = cc.Save(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.UpdatedAt, updatedAt)
	err = uow.Commit()
	assert.NoError(t, err)
	assert.Equal(t, first.Name, "repository uow update test")
	assert.NotEqual(t, first.UpdatedAt, updatedAt)
	DeleteMockTest(conn)
}

func TestUowRepositoryRemove(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	first := Test{}
	uow := conn.Uow()
	c := conn.Db(first, uow)
	first.Name = "repository uow remove test"
	err := c.Add(&first)
	err = c.Remove(&first)
	assert.NoError(t, err)
	assert.False(t, first.DeletedAt.Valid)
	err = uow.Commit()
	assert.NoError(t, err)
	assert.True(t, first.DeletedAt.Valid)
	DeleteMockTest(conn)
}
