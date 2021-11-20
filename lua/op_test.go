package lua_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rolancia/go-lua/lua"
	"github.com/rolancia/go-lua/lua/lualib"
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
			a := l.LocalWithName("a", lua.Num(5))
			b := l.LocalWithName("b", lua.Op3(a, lua.Op("+"), lua.Num(1)))
			c := l.LocalWithName("c", lua.Op3(lua.Num(5), lua.Op("+"), lua.Num(1)))
			d := l.LocalWithName("d", lua.Op3(a, lua.Op("+"), c))
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
			a := l.LocalWithName("a", lua.Bool(true))
			b := l.LocalWithName("b", lua.Op2(lua.Op("not"), a))
			lualib.Print(l, a, b, lua.Op2(lua.Op("not"), b))
		})
		assert.Equal(t, expected, scr)
	})
}