package lua_test

import (
	"github.com/rolancia/go-lua/lua"
	"github.com/rolancia/go-lua/lua/lualib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoop(t *testing.T) {
	t.Run("loop", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			lua.For(l, 1, 10, 1).Do(func(i lua.Variable) {
				lualib.Print(l, i)
			})
		})
		assert.Equal(t, reduceLMargin(`
for a = 1,10,1
do
	print(a)
end
`), scr)
	})

	t.Run("nested loop", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			lua.For(l, 1, 10, 1).Do(func(i lua.Variable) {
				lua.For(l, 1, 10, 1).Do(func(j lua.Variable) {
					lualib.Print(l, i, j)
				})
			})

			a := 1
			_ = a
		})
		assert.Equal(t, reduceLMargin(`
for a = 1,10,1
do
	for b = 1,10,1
	do
		print(a,b)
	end
end
`), scr)
	})

	t.Run("access array", func(t *testing.T) {
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			arr := l.LocalWithName("arr", lua.Table())
			lua.For(l, 1, 10, 1).Do(func(i lua.Variable) {
				l.Assign(lua.At(arr, i), i)
			})
			lua.For(l, 1, 10, 1).Do(func(i lua.Variable) {
				lualib.Print(l, lua.At(arr, i))
			})
		})
		assert.Equal(t, reduceLMargin(`
local arr = {}
for a = 1,10,1
do
	arr[a] = a
end
for a = 1,10,1
do
	print(arr[a])
end
`), scr)
	})
}
