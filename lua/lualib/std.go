package lualib

import (
	"fmt"
	"strings"

	"github.com/rolancia/go-lua/lua"
)

func Print(l *lua.Builder, args ...lua.Object) {
	luaArgs := make([]string, len(args))
	for i, arg := range args {
		switch arg.(type) {
		case lua.Variable:
			luaArgs[i] = arg.(lua.Variable).Name()
		default:
			if arg.Type() == "string" {
				luaArgs[i] = "\"" + arg.Value() + "\""
			} else {
				luaArgs[i] = fmt.Sprintf("%v", arg.Value())
			}
		}
	}
	l.Append([]byte(fmt.Sprintf("print(%s)", strings.Join(luaArgs, ","))))
	l.AppendLine()
}
