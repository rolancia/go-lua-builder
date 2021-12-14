package lua

import (
	"fmt"
	"strings"
)

func Call(ca Callable) Object {
	sArgs := make([]string, len(ca.Args()))
	for i, arg := range ca.Args() {
		sArgs[i] = arg.Tag()
	}
	v := fmt.Sprintf("%s(%s)", ca.Tag(), strings.Join(sArgs, ","))
	return Any(v)
}

// function
type Callable interface {
	Object
	Args() []Object
}

var _ Callable = &TypeFunc{}

func Func(name string, args ...Object) TypeFunc {
	return TypeFunc{
		N:    name,
		args: args,
	}
}

type TypeFunc struct {
	N    string
	args []Object
}

func (t TypeFunc) Type() string {
	return "function"
}

func (t TypeFunc) Tag() string {
	return t.N
}

func (t TypeFunc) Args() []Object {
	return t.args
}
