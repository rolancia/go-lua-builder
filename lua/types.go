package lua

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Object interface {
	Type() string
	Value() string
}

// nil
func Nil() TypeNil {
	return TypeNil{}
}

var _ Object = &TypeNil{}

type TypeNil struct {
	V interface{}
}

func (n TypeNil) Type() string {
	return "nil"
}

func (n TypeNil) Value() string {
	return "nil"
}

// boolean
func Bool(v bool) TypeBoolean {
	return TypeBoolean{V: v}
}

var _ Object = &TypeBoolean{}

type TypeBoolean struct {
	V bool
}

func (b TypeBoolean) Type() string {
	return "boolean"
}

func (b TypeBoolean) Value() string {
	if b.V {
		return "true"
	} else {
		return "false"
	}
}

// number
func Num(v int) TypeNumber {
	return TypeNumber{V: v}
}

var _ Object = &TypeNumber{}

type TypeNumber struct {
	V int
}

func (n TypeNumber) Type() string {
	return "number"
}

func (n TypeNumber) Value() string {
	return strconv.Itoa(n.V)
}

// string
func Str(v string) TypeString {
	return TypeString{V: v}
}

var _ Object = &TypeString{}

type TypeString struct {
	V string
}

func (s TypeString) Type() string {
	return "string"
}

func (s TypeString) Value() string {
	return fmt.Sprintf("\"%s\"", s.V)
}

// table
func At(v Variable, k Object) Variable {
	key := k.Value()
	return NewVar(fmt.Sprintf("%s[%v]", v.Name(), key), v)
}

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

func (t TypeTable) Value() string {
	strs := make([]string, len(t.V.els))
	for i := range t.V.els {
		e := t.V.els[i]
		k, v := e.k.(string), e.v
		strs[i] = fmt.Sprintf("%s = %s", k, v.Value())
	}
	return fmt.Sprintf("{%s}", strings.Join(strs, ","))
}

// array
func Array(initial ...Object) TypeArray {
	t := table{}
	for i := range initial {
		t.add(tableElement{
			k: i + 1,
			v: initial[i],
		})
	}
	return TypeArray{V: t}
}

type TypeArray struct {
	V table
}

func (t TypeArray) Type() string {
	return "array"
}

func (t TypeArray) Value() string {
	strs := make([]string, len(t.V.els))
	for i := range t.V.els {
		strs[i] = t.V.els[i].v.Value()
	}
	return fmt.Sprintf("{%s}", strings.Join(strs, ","))
}

// any
func Any(o Object) TypeAny {
	return TypeAny{
		Object: o,
		V:      o.Value(),
	}
}

type TypeAny struct {
	Object
	V string
}

func (a TypeAny) Type() string {
	return a.Object.Type()
}

func (a TypeAny) Value() string {
	return a.V
}

// function
type Callable interface {
	Variable
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

func (t TypeFunc) Name() string {
	return t.N
}

func (t TypeFunc) Type() string {
	return "function"
}

func (t TypeFunc) Value() string {
	return t.N
}

func (t TypeFunc) Args() []Object {
	return t.args
}
