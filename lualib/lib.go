package lualib

import (
	"fmt"
	"github.com/rolancia/go-lua-builder/lua"
	"strings"
)

func AppendCall(b lua.Builder, ca lua.Callable, rets ...lua.Object) {
	if len(rets) != 0 {
		sRets := make([]string, len(rets))
		for i, ret := range rets {
			sRets[i] = ret.Tag()
		}
		b.AppendStr(fmt.Sprintf("local %s = %s", strings.Join(sRets, ","), lua.Call(ca).Tag()))
	} else {
		b.AppendStr(lua.Call(ca).Tag())
	}
	b.AppendLine()
}
