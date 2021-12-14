package lua

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

/*
Object
*/
type Object interface {
	Type() string
	Tag() string
}

type NumObject Object
type StrObject Object
type BoolObject Object
type OpObject Object

/*
Var has an object with the name
*/
func newVar(n string, o Object) Var {
	return Var{
		O: o,
		N: n,
	}
}

var _ Object = &Var{}

type Var struct {
	O Object
	N string
}

func (v Var) Type() string {
	return v.O.Type()
}

func (v Var) Tag() string {
	return v.N
}

type NumVar struct{ Var }
type StrVar struct{ Var }
type BoolVar struct{ Var }
type TableVar struct{ Var }
type ArrayVar struct{ Var }

func NewNumVar(n string, v Num) NumVar {
	return NumVar{newVar(n, v)}
}

/*
Types
*/
// Number
var _ Object = Num(0.0)

type Num float64

func (n Num) Type() string {
	return "number"
}

func (n Num) Tag() string {
	if n.IsFloat() {
		return fmt.Sprintf("%.13f", n)
	} else {
		return fmt.Sprintf("%.0f", n)
	}
}

func (n Num) IsFloat() bool {
	_, f := math.Modf(float64(n))
	f = f * math.Pow10(13) / math.Pow10(13)
	return f != 0.0
}

// String
var _ Object = Str("")

type Str string

func (s Str) Type() string {
	return "string"
}

func (s Str) Tag() string {
	return fmt.Sprintf("\"%s\"", s)
}

// Boolean
var _ Object = Bool(false)

type Bool bool

func (b Bool) Type() string {
	return "boolean"
}

func (b Bool) Tag() string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

// Operator
var _ Object = OpVal("")

type OpVal string

func (o OpVal) Type() string {
	return "operator"
}

func (o OpVal) Tag() string {
	return string(o)
}

// Any
var _ Object = Any("")

type Any string

func (a Any) Type() string {
	return "any"
}

func (a Any) Tag() string {
	return string(a)
}

// Nil
var _ Object = Nil("")

type Nil string

func (n Nil) Type() string {
	return "nil"
}

func (n Nil) Tag() string {
	return "nil"
}

// base table
type table struct {
	els []tableElement
}

func (t *table) add(e tableElement) {
	t.els = append(t.els, e)
}

type tableElement struct {
	k interface{} // int or string
	v Object
}

func Table(initial map[string]Object) TypeTable {
	t := table{}
	sorted := make([]string, 0, len(initial))
	for k := range initial {
		sorted = append(sorted, k)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	for _, k := range sorted {
		t.add(tableElement{
			k: k,
			v: initial[k],
		})
	}
	return TypeTable{V: t}
}

var _ Object = &TypeTable{}

type TypeTable struct {
	V table
}

func (t TypeTable) Type() string {
	return "table"
}

func (t TypeTable) Tag() string {
	strs := make([]string, len(t.V.els))
	for i := range t.V.els {
		e := t.V.els[i]
		k, v := e.k.(string), e.v
		strs[i] = fmt.Sprintf("%s = %s", k, v.Tag())
	}
	return fmt.Sprintf("{%s}", strings.Join(strs, ","))
}

// array
var _ Object = &ArrayVal{}

func Array(initial ...Object) ArrayVal {
	t := table{}
	for i := range initial {
		t.add(tableElement{
			k: i + 1,
			v: initial[i],
		})
	}
	return ArrayVal{V: t}
}

type ArrayVal struct {
	V table
}

func (t ArrayVal) Type() string {
	return "array"
}

func (t ArrayVal) Tag() string {
	strs := make([]string, len(t.V.els))
	for i := range t.V.els {
		strs[i] = t.V.els[i].v.Tag()
	}
	return fmt.Sprintf("{%s}", strings.Join(strs, ","))
}

func At(v Object, k Object) Object {
	key := k.Tag()
	return Any(fmt.Sprintf("%s[%v]", v.Tag(), key))
}
