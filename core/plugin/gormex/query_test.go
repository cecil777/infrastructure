package gormex

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestQueryCount(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	MultipleCreateTest(conn, "query count test")
	first := Test{}
	c := conn.Db(first)
	num, err := c.Query().Count()
	fmt.Println("query count num:", num)
	assert.NoError(t, err)
	assert.Equal(t, num, int64(3))
	DeleteMockTest(conn)
}

func TestQueryOrder(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	MultipleCreateTest(conn, "query order test")
	first := Test{}
	c := conn.Db(first)
	err := c.Query().Order("name").Take(1).ToArray(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.Name, "query order test 1")
	DeleteMockTest(conn)
}

func TestQueryOrderByDesc(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	MultipleCreateTest(conn, "query order by desc test")
	first := Test{}
	c := conn.Db(first)
	err := c.Query().OrderByDesc("name").Take(1).ToArray(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.Name, "query order by desc test 3")
	DeleteMockTest(conn)
}

func TestQuerySkip(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	MultipleCreateTest(conn, "query skip test")
	first := Test{}
	c := conn.Db(first)
	err := c.Query().Skip(1).Take(1).ToArray(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.Name, "query skip test 2")
	DeleteMockTest(conn)
}

func TestQueryTake(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	MultipleCreateTest(conn, "query take test")
	first := Test{}
	c := conn.Db(first)
	var tt []Test
	err := c.Query().Take(3).ToArray(&tt)
	assert.NoError(t, err)
	assert.Equal(t, len(tt), 3)
	DeleteMockTest(conn)
}

func TestQueryToArray(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	MultipleCreateTest(conn, "query take test")
	first := Test{}
	c := conn.Db(first)
	var tt []Test
	err := c.Query().OrderByDesc("name").Take(3).ToArray(&tt)
	assert.NoError(t, err)
	assert.Equal(t, len(tt), 3)
	DeleteMockTest(conn)
}

func TestQueryWhere(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	conn := NewMock()
	MultipleCreateTest(conn, "query take test")
	first := Test{}
	c := conn.Db(first)
	var tt []Test
	err := c.Query().OrderByDesc("name").Take(3).ToArray(&tt)
	assert.NoError(t, err)
	assert.Equal(t, len(tt), 3)
	DeleteMockTest(conn)
}
