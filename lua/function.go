package lua

import (
	"fmt"
	"strings"
)

func Call(ca Callable) Object {
	sArgs := make([]string, len(ca.Args()))
	for i, arg := range ca.Args() {
		sArgs[i] = arg.Value()
	}
	v := fmt.Sprintf("%s(%s)", ca.Name(), strings.Join(sArgs, ","))
	ret := NewVar(v, Nil())
	return ret
}
