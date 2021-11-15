package lua

import "fmt"

type Loop struct {
	b     Builder
	start int
	end   int
	step  int
}

func (l Loop) Do(f func(i Variable)) {
	b := l.b
	id := string(rune(int('a') + b.NumLoop()))
	b.SetNumLoop(b.NumLoop() + 1)
	b.Append([]byte(fmt.Sprintf("for %s = %d,%d,%d", id, l.start, l.end, l.step)))
	b.AppendLine()
	b.Append([]byte("do"))
	b.AppendLine()
	withTab(b, func() {
		f(newVar(id, Num(0)))
	})
	b.Append([]byte("end"))
	b.AppendLine()
	b.SetNumLoop(b.NumLoop() - 1)
}
