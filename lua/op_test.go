package lua_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rolancia/go-lua-builder/lua"
	"github.com/rolancia/go-lua-builder/lualib"
)

func TestOp(t *testing.T) {
	t.Run("op3", func(t *testing.T) {
		expected := reduceLMargin(`
local a = 5
local b = a + 1
local c = 5 + 1
local d = a + c
print(a,b,c,d,10 + 9,a + b)
`)
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			a := l.Local(lua.Num(5), "a")
			b := l.Local(lua.Op3(a, lua.Op("+"), lua.Num(1)), "b")
			c := l.Local(lua.Op3(lua.Num(5), lua.Op("+"), lua.Num(1)), "c")
			d := l.Local(lua.Op3(a, lua.Op("+"), c), "d")
			lualib.Print(l, a, b, c, d, lua.Op3(lua.Num(10), lua.Op("+"), lua.Num(9)), lua.Op3(a, lua.Op("+"), b))
		})
		assert.Equal(t, expected, scr)
	})

	t.Run("op2", func(t *testing.T) {
		expected := reduceLMargin(`
local a = true
local b = not a
print(a,b,not b)
`)
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			a := l.Local(lua.Bool(true), "a")
			b := l.Local(lua.Op2(lua.Op("not"), a), "b")
			lualib.Print(l, a, b, lua.Op2(lua.Op("not"), b))
		})
		assert.Equal(t, expected, scr)
	})

	t.Run("not", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v := l.Local(lua.Num(1), "v")
			nv := l.Local(lua.Not(v), "nv")
			nb := l.Local(lua.Not(lua.Bool(true)), "nb")
			ns := l.Local(lua.Not(lua.Str("hello")), "ns")
			c := lua.Cond(lua.Bool(true), lua.Eq(), lua.Bool(false))
			nc := l.Local(lua.Not(c), "nc")
			_, _, _, _ = nv, nb, ns, nc
		})
		assert.Equal(t, reduceLMargin(`
local v = 1
local nv = not v
local nb = not true
local ns = not "hello"
local nc = not (true == false)
`), scr)
	})
}
