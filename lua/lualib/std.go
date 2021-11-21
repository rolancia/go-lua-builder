package lualib

import (
	"fmt"
	"strings"

	"github.com/rolancia/go-lua-builder/lua"
)

func Print(l lua.Builder, args ...lua.Object) {
	luaArgs := make([]string, len(args))
	for i, arg := range args {
		luaArgs[i] = arg.Value()
	}
	l.Append([]byte(fmt.Sprintf("print(%s)", strings.Join(luaArgs, ","))))
	l.AppendLine()
}
