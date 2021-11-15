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

type Variable interface {
	Object
	Name() string
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
func Table() TypeTable {
	v := make(map[interface{}]interface{})
	return TypeTable{V: v}
}

var _ Object = &TypeTable{}

type TypeTable struct {
	V map[interface{}]interface{}
}

func (t TypeTable) Type() string {
	return "table"
}

func (t TypeTable) Value() string {
	strs := make([]string, 0, len(t.V))
	for k := range t.V {
		strs = append(strs, fmt.Sprintf("%v", t.V[k]))
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	return fmt.Sprintf("{%s}", strings.Join(strs, ","))
}

func At(v Variable, k Object) Variable {
	key := k.Value()
	return newVar(fmt.Sprintf("%s[%v]", v.Name(), key), v)
}
