package lua

import "fmt"

type Loop struct {
	start int
	end   int
	step  int
}

func (l Loop) Do(b Builder, f func(i Variable)) {
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

func For(start, end, step int) Loop {
	l := Loop{
		start: start,
		end:   end,
		step:  step,
	}
	return l
}
