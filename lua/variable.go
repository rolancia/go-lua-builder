package lua

type Variable interface {
	Object
	Name() string
}

// var
var _ Variable = &Var{}

type Var struct {
	M Object
	N string
}

func (v Var) Type() string {
	return v.M.Type()
}

func (v Var) Value() string {
	return v.Name()
}

func (v Var) Name() string {
	return v.N
}

func NewVar(name string, m Object) Var {
	return Var{
		N: name,
		M: m,
	}
}

/*
Typed Var

Typed Var is not needed in Lua but these are defined to indicate types in go
eg. return values by function
with the typed var, function can be defined like below
func foo(args ...Object) (count NumVar, length StrVar)
*/

func NewNilVar(name string) NilVar {
	return NilVar{Var: NewVar(name, Nil())}
}

type NilVar struct{ Var }

func NewNumVar(name string, v TypeNumber) NumVar {
	return NumVar{Var: NewVar(name, v)}
}

type NumVar struct{ Var }

func NewStrVar(name string, v TypeString) StrVar {
	return StrVar{Var: NewVar(name, v)}
}

type StrVar struct{ Var }

func NewBoolVar(name string, v TypeBoolean) BoolVar {
	return BoolVar{Var: NewVar(name, v)}
}

type BoolVar struct{ Var }

func NewTableVar(name string, v TypeTable) TableVar {
	return TableVar{Var: NewVar(name, v)}
}

type TableVar struct{ Var }

func NewArrayVar(name string, v TypeArray) ArrayVar {
	return ArrayVar{Var: NewVar(name, v)}
}

type ArrayVar struct{ Var }

func NewAnyVar(name string, v TypeAny) AnyVar {
	return AnyVar{Var: NewVar(name, v)}
}

type AnyVar struct{ Var }

func NewFuncVar(name string, v TypeFunc) FuncVar {
	return FuncVar{Var: NewVar(name, v)}
}

type FuncVar struct{ Var }
