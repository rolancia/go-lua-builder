package lua

func NewLua(fn func(l *DefaultBuilder)) string {
	b := NewBuilder()
	fn(b)
	return b.String()
}
