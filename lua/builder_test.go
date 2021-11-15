package lua_test

import (
	"github.com/rolancia/go-lua/lua"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssign(t *testing.T) {
	t.Run("assignment", func(t *testing.T) {
		expected := reduceLMargin(`
local v1 = "hello"
local v2 = v1
v2 = v1
v2 = 123
`)
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.LocalWithName("v1", lua.Str("hello"))
			v2 := l.LocalWithName("v2", v1)
			l.Assign(v2, v1)
			l.Assign(v2, lua.Num(123))
		})
		assert.Equal(t, expected, scr)
	})
}
