package lualib_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rolancia/go-lua-builder/lua"
	"github.com/rolancia/go-lua-builder/lualib"
)

func TestFunctions(t *testing.T) {
	t.Run("print", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			l.Do(lua.Call(lualib.PrintCab(lua.Str("hello"))))
			lualib.Print(l, lua.Str("world"))
		})
		assert.Equal(t, reduceLMargin(`
print("hello")
print("world")
`), scr)
	})

	t.Run("tonumber", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			lualib.Print(l, lua.Call(lualib.ToNumberCab(lua.Str("12345"))))
			n := lualib.ToNumber(l, lua.Str("6789"))
			lualib.Print(l, n)
		})
		assert.Equal(t, reduceLMargin(`
print(tonumber("12345"))
local number1 = tonumber("6789")
print(number1)
`), scr)
	})
}
