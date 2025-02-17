package gormex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryQuery(t *testing.T) {
	test := Test{}
	g := NewFactory(nil)
	_, ok := g.Db(test).Query().(*query)
	assert.True(t, ok)
}

func TestRepositoryAdd(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	first := Test{}
	c := conn.Db(first)
	first.Name = "repository add test"
	err := c.Add(&first)
	assert.NoError(t, err)
	assert.NotEqual(t, first.ID, uint(0))
	assert.Equal(t, first.Name, "repository add test")
}

func TestRepositorySave(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	first := Test{}
	c := conn.Db(first)
	first.Name = "repository save test"
	err := c.Add(&first)
	assert.NoError(t, err)
	first.Name = "repository update test"
	err = c.Save(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.Name, "repository update test")
}

func TestRepositoryRemove(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	first := Test{}
	c := conn.Db(first)
	first.Name = "repository remove test"
	err := c.Add(&first)
	assert.NoError(t, err)
	assert.NotEqual(t, first.ID, uint(0))
	err = c.Remove(&first)
	assert.NoError(t, err)
	assert.True(t, first.DeletedAt.Valid)
}

func TestUowRepositoryAdd(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	first := Test{}
	uow := conn.Uow()
	c := conn.Db(first, uow)
	first.Name = "repository uow add test"
	err := c.Add(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.ID, uint(0))
	first.Name = "repository uow update test"
	err = c.Save(&first)
	assert.NoError(t, err)
	err = uow.Commit()
	assert.NoError(t, err)
	assert.NotEqual(t, first.ID, uint(0))
	assert.Equal(t, first.Name, "repository uow update test")
}

func TestUowRepositorySave(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	first := Test{}
	c := conn.Db(first)
	first.Name = "repository uow save test"
	err := c.Add(&first)
	assert.NoError(t, err)
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
}

func TestUowRepositoryRemove(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	first := Test{}
	uow := conn.Uow()
	c := conn.Db(first, uow)
	first.Name = "repository uow remove test"
	err := c.Add(&first)
	assert.NoError(t, err)
	err = c.Remove(&first)
	assert.NoError(t, err)
	assert.False(t, first.DeletedAt.Valid)
	err = uow.Commit()
	assert.NoError(t, err)
	assert.True(t, first.DeletedAt.Valid)
}
