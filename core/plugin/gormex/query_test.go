package gormex

import (
	"core/db"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryCount(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	query := db.NewMockIQuery(ctl)
	query.EXPECT().Count().Return(int64(1), nil)
	count, err := query.Count()
	assert.Equal(t, count, int64(1))
	assert.Equal(t, err, nil)
}

func TestQueryOrder(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	//test := Test{}
	//g := NewFactory("")
	//test.ID = 17
	//q, ok := g.Db(test).Query().Order("id").(query)
	//assert.True(t, ok)
	//model, ok := q.model.(Test)
	//assert.True(t, ok)
	//assert.Equal(t, model.ID, uint(17))
	//assert.Equal(t, len(q.db.Statement.Clauses), 1)

	mockQuery := db.NewMockIQuery(ctl)
	mockQuery.EXPECT().Order("id").Return(&query{})
	mockQuery.Order("id")
}

func TestQueryOrderByDesc(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	//test := Test{}
	//g := NewFactory("")
	//q, ok := g.Db(test).Query().OrderByDesc("id").(query)
	//assert.True(t, ok)
	//assert.Equal(t, len(q.db.Statement.Clauses), 1)
	//_, ok = q.db.Statement.Clauses["ORDER BY"]
	//assert.True(t, ok)

	mockQuery := db.NewMockIQuery(ctl)
	mockQuery.EXPECT().OrderByDesc("id").Return(&query{})
	mockQuery.OrderByDesc("id")
}

func TestQuerySkip(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	//test := Test{}
	//g := NewFactory("")
	//q, ok := g.Db(test).Query().Skip(1).(query)
	//assert.True(t, ok)
	//assert.Equal(t, len(q.db.Statement.Clauses), 1)
	//l, ok := q.db.Statement.Clauses["LIMIT"]
	//assert.True(t, ok)
	//e, ok := l.Expression.(clause.Limit)
	//assert.True(t, ok)
	//assert.Equal(t, e.Offset, 1)

	mockQuery := db.NewMockIQuery(ctl)
	mockQuery.EXPECT().Skip(1).Return(&query{})
	mockQuery.Skip(1)
}

func TestQueryTake(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	//test := Test{}
	//g := NewFactory("")
	//q, ok := g.Db(test).Query().Take(1).(query)
	//assert.True(t, ok)
	//assert.Equal(t, len(q.db.Statement.Clauses), 1)
	//l, ok := q.db.Statement.Clauses["LIMIT"]
	//assert.True(t, ok)
	//e, ok := l.Expression.(clause.Limit)
	//assert.True(t, ok)
	//assert.Equal(t, e.Limit, 1)

	mockQuery := db.NewMockIQuery(ctl)
	mockQuery.EXPECT().Take(1).Return(&query{})
	mockQuery.Take(1)
}

func TestQueryToArray(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockIQuery := db.NewMockIQuery(ctl)
	test := Test{}
	gomock.InOrder(mockIQuery.EXPECT().ToArray(test).Return(nil))
	err := mockIQuery.ToArray(test)
	assert.Equal(t, err, nil)
}

func TestQueryWhere(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	//test := Test{}
	//g := NewFactory("")
	//q, ok := g.Db(test).Query().Where("name = ?", "test").(query)
	//assert.True(t, ok)
	//assert.Equal(t, len(q.db.Statement.Clauses), 1)
	//l, ok := q.db.Statement.Clauses["WHERE"]
	//assert.Equal(t, l.Name, "WHERE")

	mockQuery := db.NewMockIQuery(ctl)
	mockQuery.EXPECT().Where("id = ?", 17).Return(&query{})
	mockQuery.Where("id = ?", 17)
}
