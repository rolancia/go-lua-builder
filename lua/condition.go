package lua

import (
	"fmt"
)

const (
	condOPAnd = "and"
	condOPOr  = "or"
)

func Cond(a Object, op string, b Object) Condition {
	if op == "!=" {
		op = "~="
	}
	return Condition{L: a, OP: op, R: b}
}

func And(a, b Condition) Condition {
	return Condition{L: a, R: b, OP: condOPAnd}
}

func Or(a, b Condition) Condition {
	return Condition{L: a, R: b, OP: condOPOr}
}

func True() Condition {
	return Condition{OP: "true"}
}

func False() Condition {
	return Condition{OP: "false"}
}

//
type Condition struct {
	L  interface{}
	R  interface{}
	OP string
}

func (c Condition) append(buf *[]byte) {
	if c.OP == "true" || c.OP == "false" {
		*buf = append(*buf, []byte(c.OP)...)
		return
	}
	cs := []interface{}{c.L, c.R}
	for i, clr := range cs {
		if lc, ok := clr.(Variable); ok {
			*buf = append(*buf, []byte(lc.Name())...)
		} else if lc, ok := clr.(Object); ok {
			var v string
			if lc.Type() == "string" {
				v = fmt.Sprintf("\"%s\"", lc.Value())
			} else {
				v = lc.Value()
			}
			*buf = append(*buf, []byte(v)...)
		} else if in, ok := clr.(Condition); ok {
			*buf = append(*buf, '(')
			in.append(buf)
			*buf = append(*buf, ')')
		} else {
			panic(fmt.Errorf("invalid condition type"))
		}

		if i == 0 {
			*buf = append(*buf, fmt.Sprintf(" %s ", c.OP)...)
		}
	}
}

func (c Condition) ToBytes() []byte {
	buf := make([]byte, 0, 8)
	c.append(&buf)
	return buf
}

type condBuilder struct{ b *Builder }

func beginIf(b *Builder, c Condition) IfThen {
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
