package cor

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testContext struct {
	count int
}

type testHandler struct {
	IHandler

	count int
	ctx   *testContext
}

func (m *testHandler) Handle() error {
	if m.count < 0 {
		return errors.New("err")
	}

	if m.count == 0 {
		m.Break()
	}

	m.ctx.count = m.ctx.count + m.count
	return m.IHandler.Handle()
}

func Test_handler_Break(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		ctx := new(testContext)
		h := &testHandler{
			IHandler: New(),
			ctx:      ctx,
		}
		h.SetNext(&testHandler{
			IHandler: New(),
			count:    2,
			ctx:      ctx,
		})
		err := h.Handle()
		assert.NoError(t, err)
		assert.Equal(t, ctx.count, 0)
	})
	t.Run("second", func(t *testing.T) {
		ctx := new(testContext)
		h := &testHandler{
			IHandler: New(),
			count:    1,
			ctx:      ctx,
		}
		h.SetNext(&testHandler{
			IHandler: New(),
			ctx:      ctx,
		}).SetNext(&testHandler{
			IHandler: New(),
			count:    3,
			ctx:      ctx,
		})
		err := h.Handle()
		assert.NoError(t, err)
		assert.Equal(t, ctx.count, 1)
	})
	t.Run("ok", func(t *testing.T) {
		ctx := new(testContext)
		h := &testHandler{
			IHandler: New(),
			count:    1,
			ctx:      ctx,
		}
		h.SetNext(&testHandler{
			IHandler: New(),
			count:    2,
			ctx:      ctx,
		}).SetNext(&testHandler{
			IHandler: New(),
			count:    3,
			ctx:      ctx,
		})
		err := h.Handle()
		assert.NoError(t, err)
		assert.Equal(t, ctx.count, 6)
	})
}
