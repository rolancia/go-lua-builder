package lua_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rolancia/go-lua-builder/lua"
)

func TestTypes(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.Local(lua.Nil(""), "v1")
			l.Assign(v1, lua.Nil(""))
		})
		assert.Equal(t, reduceLMargin(`
local v1 = nil
v1 = nil
`), scr)
	})

	t.Run("boolean", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.Local(lua.Bool(true), "v1")
			v2 := l.Local(lua.Bool(false), "v2")
			l.Assign(v1, lua.Bool(false))
			l.Assign(v2, lua.Bool(true))
		})
		assert.Equal(t, reduceLMargin(`
local v1 = true
local v2 = false
v1 = false
v2 = true
`), scr)
	})

	t.Run("number", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.Local(lua.Num(1), "v1")
			v2 := l.Local(lua.Num(2), "v2")
			l.Assign(v1, v2)
			l.Assign(v2, lua.Num(3))
		})
		assert.Equal(t, reduceLMargin(`
local v1 = 1
local v2 = 2
v1 = v2
v2 = 3
`), scr)
	})

	t.Run("string", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.Local(lua.Str("hello"), "v1")
			v2 := l.Local(lua.Str("world"), "v2")
			l.Assign(v1, v2)
			l.Assign(v2, lua.Str("!"))
		})
		assert.Equal(t, reduceLMargin(`
local v1 = "hello"
local v2 = "world"
v1 = v2
v2 = "!"
`), scr)
	})

	t.Run("array", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			arr1 := l.Local(lua.Array(), "arr1")
			arr2 := l.Local(lua.Array(lua.Num(1), lua.Num(2), lua.Str("hello"), lua.Str("world")), "arr2")
			_ = arr2
			l.Assign(lua.At(arr1, lua.Num(1)), lua.Num(2021))
		})
		assert.Equal(t, reduceLMargin(`
local arr1 = {}
local arr2 = {1,2,"hello","world"}
arr1[1] = 2021
`), scr)
	})

	t.Run("table", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			t1 := l.Local(lua.Table(nil), "t1")
			t2 := l.Local(lua.Table(map[string]lua.Object{
				"a": lua.Str("this is a"),
				"b": lua.Num(2),
			}), "t2")
			_ = t2
			l.Assign(lua.At(t1, lua.Str("a")), lua.Str("you are not a"))
		})
		assert.Equal(t, reduceLMargin(`
local t1 = {}
local t2 = {a = "this is a",b = 2}
t1["a"] = "you are not a"
`), scr)
	})

	t.Run("any", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.Local(lua.Num(1), "v1")
			v2 := l.Local(lua.Str("hello"), "v2")
			_, _ = v1, v2
		})
		assert.Equal(t, reduceLMargin(`
local v1 = 1
local v2 = "hello"
`), scr)
	})
}
