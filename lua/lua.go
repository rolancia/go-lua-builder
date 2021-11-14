package lua

func NewLua(fn func(l *Builder)) string {
	b := &Builder{}
	fn(b)
	return b.String()
}
