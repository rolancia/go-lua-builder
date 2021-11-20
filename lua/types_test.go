package lua_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rolancia/go-lua/lua"
)

func TestTypes(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.LocalWithName("v1", lua.Nil())
			l.Assign(v1, lua.Nil())
		})
		assert.Equal(t, reduceLMargin(`
local v1 = nil
v1 = nil
`), scr)
	})

	t.Run("boolean", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.LocalWithName("v1", lua.Bool(true))
			v2 := l.LocalWithName("v2", lua.Bool(false))
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
			v1 := l.LocalWithName("v1", lua.Num(1))
			v2 := l.LocalWithName("v2", lua.Num(2))
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
			v1 := l.LocalWithName("v1", lua.Str("hello"))
			v2 := l.LocalWithName("v2", lua.Str("world"))
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
			arr1 := l.LocalWithName("arr1", lua.Array())
			arr2 := l.LocalWithName("arr2", lua.Array(lua.Num(1), lua.Num(2), lua.Str("hello"), lua.Str("world")))
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
			t1 := l.LocalWithName("t1", lua.Table(nil))
			t2 := l.LocalWithName("t2", lua.Table(map[string]lua.Object{
				"a": lua.Str("this is a"),
				"b": lua.Num(2),
			}))
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
			v1 := l.LocalWithName("v1", lua.Num(1))
			v11 := l.LocalWithName("v11", lua.Any(v1))
			v111 := l.LocalWithName("v111", lua.Any(lua.Num(111)))
			v2 := l.LocalWithName("v2", lua.Str("hello"))
			v22 := l.LocalWithName("v22", lua.Any(v2))
			v222 := l.LocalWithName("v222", lua.Any(lua.Str("world")))
			_, _, _, _ = v11, v111, v22, v222
		})
		assert.Equal(t, reduceLMargin(`
local v1 = 1
local v11 = v1
local v111 = 111
local v2 = "hello"
local v22 = v2
local v222 = "world"
`), scr)
	})
}
