package lua

import (
	"fmt"
	"strings"
)

type Operator interface {
	Op() string
}

var allowedOps = map[string]struct{}{
	// assign
	"=": {},
	// relational
	"==": {}, "~=": {}, ">": {}, "<": {}, ">=": {}, "<=": {},
	// logical
	"and": {}, "or": {}, "not": {},
	// arithmetic
	"+": {}, "-": {}, "*": {}, "/": {}, "%": {}, "^": {},
}

var _ Operator = &BasicOperator{}

type BasicOperator struct {
	op string
}

func (op BasicOperator) Op() string {
	return op.op
}

func Op(op string) Operator {
	if _, ok := allowedOps[op]; !ok {
		panic(fmt.Errorf("op %s is not allowed", op))
	}
	return BasicOperator{op: op}
}

func Op2(op Operator, opr Object) OpObject {
	return OpVal(fmt.Sprintf("%s %s", op.Op(), opr.Tag()))
}

func Op3(l Object, op Operator, r Object) OpObject {
	return OpVal(fmt.Sprintf("%s %s %s", l.Tag(), op.Op(), r.Tag()))
}

func And() Operator {
	return newBasicOp("and")
}

func Or() Operator {
	return newBasicOp("or")
}

func Eq() Operator {
	return newBasicOp("==")
}

func Gt() Operator {
	return newBasicOp(">")
}

func Gte() Operator {
	return newBasicOp(">=")
}

func Lt() Operator {
	return newBasicOp("<")
}

func Lte() Operator {
	return newBasicOp("<=")
}

func Ne() Operator {
	return newBasicOp("~=")
}

func Not(o Object) Object {
	var v Any
	if c, ok := o.(Condition); ok {
		// wrap it if condition
		if !strings.HasPrefix(c.Tag(), "(") {
			v = Any(fmt.Sprintf("not (%s)", o.Tag()))
		}
	} else {
		v = Any(fmt.Sprintf("not %s", o.Tag()))
	}
	return v
}

func Sum(a, b Object) Object {
	return Op3(a, newBasicOp("+"), b)
}

func Sub(a, b Object) Object {
	return Op3(a, newBasicOp("-"), b)
}

func Mul(a, b Object) Object {
	return Op3(a, newBasicOp("*"), b)
}

func Div(a, b Object) Object {
	return Op3(a, newBasicOp("/"), b)
}

func Mod(a, b Object) Object {
	return Op3(a, newBasicOp("%"), b)
}

func Pow(a, b Object) Object {
	return Op3(a, newBasicOp("^"), b)
}

func newBasicOp(op string) Operator {
	return BasicOperator{op: op}
}
