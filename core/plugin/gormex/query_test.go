package gormex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryCount(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	MultipleCreateTest(conn, "query count test")
	first := Test{}
	c := conn.Db(first)
	num, err := c.Query().Count()
	assert.NoError(t, err)
	assert.Equal(t, num, int64(3))
}

func TestQueryOrder(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	MultipleCreateTest(conn, "query order test")
	first := Test{}
	c := conn.Db(first)
	err := c.Query().Order("name").Take(1).ToArray(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.Name, "query order test 1")
}

func TestQueryOrderByDesc(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	MultipleCreateTest(conn, "query order by desc test")
	first := Test{}
	c := conn.Db(first)
	err := c.Query().OrderByDesc("name").Take(1).ToArray(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.Name, "query order by desc test 3")
}

func TestQuerySkip(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	MultipleCreateTest(conn, "query skip test")
	first := Test{}
	c := conn.Db(first)
	err := c.Query().Skip(1).Take(1).ToArray(&first)
	assert.NoError(t, err)
	assert.Equal(t, first.Name, "query skip test 2")
}

func TestQueryTake(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	MultipleCreateTest(conn, "query take test")
	first := Test{}
	c := conn.Db(first)
	var tt []Test
	err := c.Query().Take(3).ToArray(&tt)
	assert.NoError(t, err)
	assert.Equal(t, len(tt), 3)
}

func TestQueryToArray(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	MultipleCreateTest(conn, "query take test")
	first := Test{}
	c := conn.Db(first)
	var tt []Test
	err := c.Query().OrderByDesc("name").Take(3).ToArray(&tt)
	assert.NoError(t, err)
	assert.Equal(t, len(tt), 3)
}

func TestQueryWhere(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	MultipleCreateTest(conn, "query take test")
	first := Test{}
	c := conn.Db(first)
	var tt []Test
	err := c.Query().OrderByDesc("name").Take(3).ToArray(&tt)
	assert.NoError(t, err)
	assert.Equal(t, len(tt), 3)
}
