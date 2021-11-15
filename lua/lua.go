package lua

func NewLua(fn func(l *DefaultBuilder)) string {
	b := &DefaultBuilder{}
	fn(b)
	return b.String()
}
