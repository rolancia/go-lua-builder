package lualib

import (
	"fmt"
	"github.com/rolancia/go-lua-builder/lua"
	"strings"
)

func AppendCall(b lua.Builder, ca lua.Callable, rets ...lua.Variable) {
	if len(rets) != 0 {
		sRets := make([]string, len(rets))
		for i, ret := range rets {
			sRets[i] = ret.Name()
		}
		b.AppendStr(fmt.Sprintf("local %s = %s", strings.Join(sRets, ","), lua.Call(ca).Value()))
	} else {
		b.AppendStr(lua.Call(ca).Value())
	}
	b.AppendLine()
}
