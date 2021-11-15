package lua

import (
	"fmt"
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
}

var _ Builder = &DefaultBuilder{}

func NewBuilder() *DefaultBuilder {
	b := &DefaultBuilder{}
	b.addr = b
	return b
}

type DefaultBuilder struct {
	addr *DefaultBuilder
	buf  []byte

	numTab  int
	numLoop int
}

func (b *DefaultBuilder) Local(m Object) Var {
	b.copyCheck()
	return b.local(randName(m.Type()), m)
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
	b.Append([]byte(dst.Name()))
	b.AppendNoTab([]byte(" = "))
	r := src.Value()
	b.AppendNoTab([]byte(r))
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

func (b *DefaultBuilder) Return(ret Object) {
	b.copyCheck()
	b.Append([]byte(fmt.Sprintf("return %s", ret.Value())))
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
