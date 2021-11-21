package lua_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rolancia/go-lua-builder/lua"
)

func TestLocal(t *testing.T) {
	t.Run("local with auto gen name", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			_ = l.Local(lua.Str("hello"))
			_ = l.Local(lua.Str("world"))
			_ = l.Local(lua.Any(lua.Str("any1")))
			_ = l.Local(lua.Any(lua.Str("any2")))
			_ = l.Local(lua.Num(1))
			_ = l.Local(lua.Num(2))
			_ = l.Local(lua.Any(lua.Num(10)))
			_ = l.Local(lua.Any(lua.Num(20)))
			_ = l.Local(lua.Bool(true))
			_ = l.Local(lua.Bool(false))
		})
		assert.Equal(t, reduceLMargin(`
local string1 = "hello"
local string2 = "world"
local string3 = "any1"
local string4 = "any2"
local number1 = 1
local number2 = 2
local number3 = 10
local number4 = 20
local boolean1 = true
local boolean2 = false
`), scr)
	})
}

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

func TestReturn(t *testing.T) {
	t.Run("just return", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			l.Return(lua.Str("bye"))
		})
		assert.Equal(t, reduceLMargin(`
return "bye"
`), scr)
	})

	t.Run("multiple return", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v := l.LocalWithName("v", lua.Num(100))
			l.Return(lua.Str("hello"), lua.Str("world"), v)
		})
		assert.Equal(t, reduceLMargin(`
local v = 100
return "hello","world",v
`), scr)
	})
}

func TestDo(t *testing.T) {
	t.Run("do", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			l.Do(lua.Op2(lua.Op("not"), lua.Num(1)))
			l.Do(lua.Op3(lua.Num(1), lua.Op("=="), lua.Num(1)))

			// assign
			v1 := l.LocalWithName("v1", lua.Nil())
			l.Do(lua.Op3(v1, lua.Op("="), lua.Str("hi")))
		})
		assert.Equal(t, reduceLMargin(`
not 1
1 == 1
local v1 = nil
v1 = "hi"
`), scr)
	})
}
