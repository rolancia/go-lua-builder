package lua_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rolancia/go-lua/lua"
	"github.com/rolancia/go-lua/lua/lualib"
)

func TestCondition(t *testing.T) {
	t.Run("condition to lua", func(t *testing.T) {
		expected := reduceMargin(`
(((v1 == v2) and (99 < 999)) or ((v1 >= v2) or (v2 >= 999))) or (true)
`)
		_ = lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.LocalWithName("v1", lua.Str("hi1"))
			v2 := l.LocalWithName("v2", lua.Str("hi2"))
			cond := lua.Or(
				lua.Or(
					lua.And(
						lua.Cond(v1, lua.Op("=="), v2),
						lua.Cond(lua.Num(99), lua.Op("<"), lua.Num(999))),
					lua.Or(
						lua.Cond(v1, lua.Op(">="), v2),
						lua.Cond(v2, lua.Op(">="), lua.Num(999))),
				),
				lua.True(),
			)
			scr := string(cond.ToBytes())
			assert.Equal(t, expected, scr)
		})
	})

	t.Run("if condition", func(t *testing.T) {
		expected := reduceLMargin(`
local v1 = "hi1"
local v2 = "hi2"
if (v1 == v2) and (1 < 10) then
	print(v1,v2,"case1")
elseif v1 > v2 then
	print(v1,v2,"case2")
elseif v1 < v2 then
	print(v1,v2,"case3")
else
	print(v1,v2,"case4")
end
`)
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.LocalWithName("v1", lua.Str("hi1"))
			v2 := l.LocalWithName("v2", lua.Str("hi2"))
			l.If(lua.And(lua.Cond(v1, lua.Op("=="), v2), lua.Cond(lua.Num(1), lua.Op("<"), lua.Num(10)))).Then(func() {
				lualib.Print(l, v1, v2, lua.Str("case1"))
			}).ElseIf(lua.Cond(v1, lua.Op(">"), v2)).Then(func() {
				lualib.Print(l, v1, v2, lua.Str("case2"))
			}).ElseIf(lua.Cond(v1, lua.Op("<"), v2)).Then(func() {
				lualib.Print(l, v1, v2, lua.Str("case3"))
			}).Else(func() {
				lualib.Print(l, v1, v2, lua.Str("case4"))
			}).End()
		})
		assert.Equal(t, expected, scr)
	})

	t.Run("nested if condition", func(t *testing.T) {
		expected := reduceLMargin(`
local v1 = "hi1"
local v2 = "hi2"
if v1 ~= v2 then
	if v1 == v2 then
		print("case1")
	end
else
	if v1 == v2 then
		print("case2")
		if v1 < v2 then
			print("case2-1")
		else
			print("case2-2")
		end
	elseif true then
		print("case3")
	end
end
`)
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			v1 := l.LocalWithName("v1", lua.Str("hi1"))
			v2 := l.LocalWithName("v2", lua.Str("hi2"))
			l.If(lua.Cond(v1, lua.Op("~="), v2)).Then(func() {
				l.If(lua.Cond(v1, lua.Op("=="), v2)).Then(func() {
					lualib.Print(l, lua.Str("case1"))
				}).End()
			}).Else(func() {
				l.If(lua.Cond(v1, lua.Op("=="), v2)).Then(func() {
					lualib.Print(l, lua.Str("case2"))
					l.If(lua.Cond(v1, lua.Op("<"), v2)).Then(func() {
						lualib.Print(l, lua.Str("case2-1"))
					}).Else(func() {
						lualib.Print(l, lua.Str("case2-2"))
					}).End()
				}).ElseIf(lua.True()).Then(func() {
					lualib.Print(l, lua.Str("case3"))
				}).End()
			}).End()
		})
		assert.Equal(t, expected, scr)
	})

	t.Run("with operator", func(t *testing.T) {
		expected := reduceLMargin(`
local a = false
if not a then
	print("hello")
else
	print("bye")
end
`)
		scr := lua.NewLua(func(l *lua.DefaultBuilder) {
			a := l.LocalWithName("a", lua.Bool(false))
			l.If(lua.Cond1(lua.Op2(lua.Op("not"), a))).Then(func() {
				lualib.Print(l, lua.Str("hello"))
			}).Else(func() {
				lualib.Print(l, lua.Str("bye"))
			}).End()
		})
		assert.Equal(t, expected, scr)
	})
}
