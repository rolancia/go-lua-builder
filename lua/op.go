package lua

import (
	"fmt"
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

func Op2(op Operator, opr Object) Object {
	v := newVar(fmt.Sprintf("%s %s", op.Op(), opr.Value()), Nil())
	return v
}

func Op3(l Object, op Operator, r Object) Object {
	v := newVar(fmt.Sprintf("%s %s %s", l.Value(), op.Op(), r.Value()), Nil())
	return v
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

func newBasicOp(op string) Operator {
	return BasicOperator{op: op}
}
