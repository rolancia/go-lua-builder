package lua

import (
	"fmt"
)

func Cond(a Object, op Operator, b Object) Condition {
	return Condition{L: a, OP: op, R: b}
}

func Cond1(o Object) Condition {
	return Condition{
		OP: condSingleOperator{b: o.Value()},
	}
}

func And(a, b Condition) Condition {
	return Condition{L: a, R: b, OP: Op("and")}
}

func Or(a, b Condition) Condition {
	return Condition{L: a, R: b, OP: Op("or")}
}

var _ Operator = &condSingleOperator{}

type condSingleOperator struct {
	b string
}

func (op condSingleOperator) Op() string {
	return op.b
}

func True() Condition {
	return Condition{OP: condSingleOperator{b: "true"}}
}

func False() Condition {
	return Condition{OP: condSingleOperator{b: "false"}}
}

//
type Condition struct {
	L  interface{}
	R  interface{}
	OP Operator
}

func (c Condition) append(buf *[]byte) {
	if c.L == nil && c.R == nil {
		*buf = append(*buf, []byte(c.OP.Op())...)
		return
	}
	cs := []interface{}{c.L, c.R}
	for i, clr := range cs {
		if lc, ok := clr.(Object); ok {
			*buf = append(*buf, lc.Value()...)
		} else if in, ok := clr.(Condition); ok {
			*buf = append(*buf, '(')
			in.append(buf)
			*buf = append(*buf, ')')
		} else {
			panic(fmt.Errorf("invalid condition type"))
		}

		if i == 0 {
			*buf = append(*buf, fmt.Sprintf(" %s ", c.OP.Op())...)
		}
	}
}

func (c Condition) ToBytes() []byte {
	buf := make([]byte, 0, 8)
	c.append(&buf)
	return buf
}

type condBuilder struct{ b *DefaultBuilder }

func beginIf(b *DefaultBuilder, c Condition) IfThen {
	b.Append([]byte("if "))
	b.AppendNoTab(c.ToBytes())
	return IfThen{b: b}
}

type IfThen condBuilder

func (t IfThen) Then(f func()) Elif {
	t.b.AppendNoTab([]byte(" then"))
	t.b.AppendLine()
	withTab(t.b, f)
	return Elif{b: t.b}
}

type ElifThen condBuilder

func (e ElifThen) Then(f func()) Elif {
	e.b.AppendNoTab([]byte(" then"))
	e.b.AppendLine()
	withTab(e.b, f)
	return Elif{b: e.b}
}

type El condBuilder

func (e El) End() {
	e.b.Append([]byte("end"))
	e.b.AppendLine()
}

type Elif condBuilder

func (e Elif) ElseIf(c Condition) ElifThen {
	e.b.Append([]byte("elseif "))
	e.b.AppendNoTab(c.ToBytes())
	return ElifThen{b: e.b}
}

func (e Elif) Else(f func()) El {
	e.b.Append([]byte("else"))
	e.b.AppendLine()
	withTab(e.b, f)
	return El{b: e.b}
}

func (e Elif) End() {
	e.b.Append([]byte("end"))
	e.b.AppendLine()
}
