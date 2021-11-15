package lua

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

func newVar(name string, m Object) Var {
	return Var{
		N: name,
		M: m,
	}
}
