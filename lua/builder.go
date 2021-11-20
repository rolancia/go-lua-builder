package lua

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"
)

type Builder interface {
	Append(b []byte)
	AppendLine()
	ApplyTabs()

	NumTab() int
	NumLoop() int
	SetNumTab(n int)
	SetNumLoop(n int)

	// NextName must be goroutine safe
	NextVariableName(ty string) string
}

var _ Builder = &DefaultBuilder{}

type DefaultBuilder struct {
	addr *DefaultBuilder
	buf  []byte

	numTab  int
	numLoop int
	counter varNameCounter
}

func (b *DefaultBuilder) Local(m Object) Var {
	b.copyCheck()
	return b.local(b.NextVariableName(m.Type()), m)
}

func (b *DefaultBuilder) LocalWithName(name string, m Object) Var {
	b.copyCheck()
	return b.local(name, m)
}

func (b *DefaultBuilder) local(name string, m Object) Var {
	b.Append([]byte(fmt.Sprintf("local %s = %s", name, m.Value())))
	b.AppendLine()
	return newVar(name, m)
}

func (b *DefaultBuilder) Assign(dst Variable, src Object) {
	b.copyCheck()
	b.Do(Op3(dst, Op("="), src))
}

func (b *DefaultBuilder) Do(v Object) {
	b.copyCheck()
	b.Append([]byte(v.Value()))
	b.AppendLine()
}

func (b *DefaultBuilder) If(c Condition) IfThen {
	b.copyCheck()
	return beginIf(b, c)
}

func (b *DefaultBuilder) For(start, end, step int) Loop {
	b.copyCheck()
	l := Loop{
		b:     b,
		start: start,
		end:   end,
		step:  step,
	}
	return l
}

func (b *DefaultBuilder) Return(rets ...Object) {
	b.copyCheck()
	strs := make([]string, len(rets))
	for i := range rets {
		strs[i] = rets[i].Value()
	}
	b.Append([]byte(fmt.Sprintf("return %s", strings.Join(strs, ","))))
	b.AppendLine()
}

func (b *DefaultBuilder) String() string {
	b.copyCheck()
	return *(*string)(unsafe.Pointer(&b.buf))
}

func (b *DefaultBuilder) Reset() {
	b.copyCheck()
	b.addr = nil
	b.buf = nil
}

func (b *DefaultBuilder) Append(bs []byte) {
	b.copyCheck()
	b.ApplyTabs()
	b.AppendNoTab(bs)
}

func (b *DefaultBuilder) AppendNoTab(bs []byte) {
	b.copyCheck()
	b.buf = append(b.buf, bs...)
}

func (b *DefaultBuilder) AppendLine() {
	b.copyCheck()
	b.buf = append(b.buf, '\n')
}

func (b *DefaultBuilder) ApplyTabs() {
	b.copyCheck()
	for i := 0; i < b.numTab; i++ {
		b.buf = append(b.buf, '\t')
	}
}

func (b *DefaultBuilder) NumTab() int {
	b.copyCheck()
	return b.numTab
}

func (b *DefaultBuilder) SetNumTab(n int) {
	b.copyCheck()
	b.numTab = n
}

func (b *DefaultBuilder) NumLoop() int {
	b.copyCheck()
	return b.numLoop
}

func (b *DefaultBuilder) SetNumLoop(n int) {
	b.copyCheck()
	b.numLoop = n
}

func (b *DefaultBuilder) NextVariableName(ty string) string {
	i := b.counter.next(ty)
	return fmt.Sprintf("%s%d", ty, i)
}

// var name counter
type varNameCounter struct {
	m map[string]int
	l sync.Mutex
}

func (c *varNameCounter) next(ty string) (new int) {
	if c.m == nil {
		c.m = make(map[string]int)
	}
	c.l.Lock()
	defer c.l.Unlock()

	if _, ok := c.m[ty]; ok {
		c.m[ty]++
	} else {
		c.m[ty] = 1
	}
	return c.m[ty]
}

//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func (b *DefaultBuilder) copyCheck() {
	if b.addr == nil {
		b.addr = (*DefaultBuilder)(noescape(unsafe.Pointer(b)))
	} else if b.addr != b {
		panic("lua-builder: illegal use of non-zero DefaultBuilder copied by value")
	}
}
