package lua_test

import (
	"github.com/rolancia/go-lua/lua"
	"github.com/stretchr/testify/assert"
	"testing"
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

	t.Run("table", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.LocalWithName("v1", lua.Table())
			v2 := l.LocalWithName("v2", lua.Table())
			l.Assign(v1, v2)
			e := l.LocalWithName("e", lua.Nil())
			e1 := lua.At(v1, lua.Num(1))
			l.Assign(e, e1)
			e2 := lua.At(v2, lua.Str("key1"))
			l.Assign(e, e2)
			key := l.LocalWithName("key", lua.Str("somekey"))
			e3 := lua.At(v2, key)
			l.Assign(e, e3)
			l.Assign(e3, lua.Str("hello"))
			l.Assign(e3, e2)
		})
		assert.Equal(t, reduceLMargin(`
local v1 = {}
local v2 = {}
v1 = v2
local e = nil
e = v1[1]
e = v2["key1"]
local key = "somekey"
e = v2[key]
v2[key] = "hello"
v2[key] = v2["key1"]
`), scr)
	})
}
