package lualib

import (
	"github.com/rolancia/go-lua-builder/lua"
)

func PrintCab(args ...lua.Object) lua.Callable {
	return lua.Func("print", args...)
}

func Print(l lua.Builder, args ...lua.Object) {
	ca := PrintCab(args...)
	AppendCall(l, ca)
}

/*
local
*/
func ToNumberCab(arg lua.Object) lua.Callable {
	return lua.Func("tonumber", arg)
}

func ToNumber(l lua.Builder, arg lua.NumObject) lua.NumVar {
	ca := ToNumberCab(arg)
	ret := lua.Num(0)
	retVar := lua.NewNumVar(l.NextVariableName(ret.Type()), ret)
	AppendCall(l, ca, retVar)
	return retVar
}
