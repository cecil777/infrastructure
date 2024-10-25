package gormex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitOfWork(t *testing.T) {
	conn := NewMock()
	defer DeleteMockTest(conn)
	//先添加2条测试数据
	first := Test{}
	c := conn.Db(first)
	first.Name = "uow test 1"
	err := c.Add(&first)
	assert.NoError(t, err)
	two := Test{}
	two.Name = "uow test 2"
	err = c.Add(&two)
	assert.NoError(t, err)
	updatedAt := first.UpdatedAt
	//开启事务
	uow := conn.Uow()
	uowConn := conn.Db(first, uow)
	first.Name = "uow update test 1"
	//更新数据
	err = uowConn.Save(&first)
	assert.NoError(t, err)
	//断言更新时间不变
	assert.Equal(t, first.UpdatedAt, updatedAt)
	three := Test{}
	three.Name = "uow test 3"
	//添加新数据
	err = uowConn.Add(&three)
	assert.NoError(t, err)
	//断言ID为空
	assert.Empty(t, three.ID)
	//删除数据
	err = uowConn.Remove(&two)
	assert.NoError(t, err)
	//断言delete_at为null
	assert.False(t, two.DeletedAt.Valid)
	//事务提交
	err = uow.Commit()
	assert.NoError(t, err)
	//断言name、updated_at变更
	assert.Equal(t, first.Name, "uow update test 1")
	assert.NotEqual(t, first.UpdatedAt, updatedAt)
	assert.True(t, two.DeletedAt.Valid)
	count, err := conn.Db(Test{}).Query().Count()
	assert.NoError(t, err)
	//断言匹配
	assert.Equal(t, count, int64(2))
}
