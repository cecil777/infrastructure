package ioc

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cecil777/infrastructure/core/reflectex"

	"github.com/stretchr/testify/assert"
)

type iInterface interface {
	Test()
}

type derive struct{}

func (m derive) Test() {
	fmt.Println("set test")
}

type myTest struct {
	One iInterface `inject:""`
}

func Test_Get(t *testing.T) {
	defer func() {
		assert.Nil(
			t,
			recover(),
		)
	}()

	ct := getInterfaceType(
		(*iInterface)(nil),
	)
	instanceValues[ct] = reflect.ValueOf(1)
	defer delete(instanceValues, ct)

	res := Get(ct)
	assert.Equal(t, res, 1)
}

func Test_Get_无效类型(t *testing.T) {
	ct := getInterfaceType(
		(*iInterface)(nil),
	)
	defer func() {
		rv := recover()
		assert.NotNil(t, rv)

		err, ok := rv.(error)
		assert.True(t, ok)
		assert.Equal(
			t,
			err,
			fmt.Errorf(invalidTypeFormat, ct),
		)
	}()

	Get(ct)
}

func Test_Has(t *testing.T) {
	ct := getInterfaceType(
		(*iInterface)(nil),
	)
	instanceValues[ct] = reflect.ValueOf(1)
	defer delete(instanceValues, ct)

	assert.True(
		t,
		Has(ct),
	)
}

func Test_Inject(t *testing.T) {
	it := getInterfaceType(
		(*iInterface)(nil),
	)
	instanceValues[it] = reflect.ValueOf(
		new(derive),
	)

	var m myTest
	Inject(&m, func(v reflect.Value) reflect.Value {
		return v
	})

	assert.Equal(
		t,
		m.One,
		instanceValues[it].Interface(),
	)
}

func Test_Remove(t *testing.T) {
	it := getInterfaceType(
		(*iInterface)(nil),
	)
	defer func() {
		assert.Nil(
			t,
			recover(),
		)
	}()

	Remove(it)
}

func Test_Set(t *testing.T) {
	defer func() {
		assert.Nil(
			t,
			recover(),
		)
	}()

	ct := reflectex.InterfaceTypeOf(
		(*iInterface)(nil),
	)
	defer delete(instanceValues, ct)

	Set(
		ct,
		new(derive),
	)
}

func Test_Set_非接口类型(t *testing.T) {
	it := reflect.TypeOf(1)
	defer func() {
		rv := recover()
		assert.NotNil(t, rv)

		err, ok := rv.(error)
		assert.True(t, ok)
		assert.Equal(
			t,
			err,
			fmt.Errorf(notInterfaceTypeFormat, it),
		)
	}()
	Set(
		it,
		new(derive),
	)
}

func Test_Set_没有继承(t *testing.T) {
	it := getInterfaceType(
		(*iInterface)(nil),
	)
	v := ""
	defer func() {
		rv := recover()
		assert.NotNil(t, rv)

		err, ok := rv.(error)
		assert.True(t, ok)
		assert.Equal(
			t,
			err,
			fmt.Errorf(notImplementsFormat, v, it),
		)
	}()
	Set(it, v)
}

func Test_Set_对象(t *testing.T) {
	defer func() {
		assert.Nil(
			t,
			recover(),
		)
	}()

	ct := reflectex.InterfaceTypeOf((*iInterface)(nil))
	defer delete(instanceValues, ct)

	Set(
		(*iInterface)(nil),
		new(derive),
	)

	_, ok := instanceValues[ct]
	assert.True(t, ok)
}
