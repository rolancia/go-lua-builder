package lua

import (
	"fmt"
	"unsafe"
)

type Builder struct {
	addr *Builder
	buf  []byte

	numTab  int
	numLoop int
}

func (b *Builder) Local(m Object) Var {
	return b.local(randName(m.Type()), m)
}

func (b *Builder) LocalWithName(name string, m Object) Var {
	return b.local(name, m)
}

func (b *Builder) local(name string, m Object) Var {
	if v, ok := m.(Variable); ok {
		b.buf = append(b.buf, fmt.Sprintf("local %s = %s", name, v.Name())...)
	} else if m.Type() == "string" {
		b.buf = append(b.buf, fmt.Sprintf("local %s = \"%s\"", name, m.Value())...)
	} else {
		b.buf = append(b.buf, fmt.Sprintf("local %s = %s", name, m.Value())...)
	}
	b.AppendLine()
	return newVar(name, m)
}

func (b *Builder) Assign(dst Variable, src Object) {
	b.Append([]byte(dst.Name()))
	b.AppendNoTab([]byte(" = "))
	var r string
	if v, ok := src.(Variable); ok {
		r = v.Name()
	} else if src.Type() == "string" {
		r = fmt.Sprintf("\"%s\"", src.Value())
	} else {
		r = src.Value()
	}
	b.AppendNoTab([]byte(r))
	b.AppendLine()
}

func (b *Builder) If(c Condition) IfThen {
	return beginIf(b, c)
}

func (b *Builder) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}

func (b *Builder) Reset() {
	b.addr = nil
	b.buf = nil
}

func (b *Builder) Append(bs []byte) {
	b.ApplyTabs()
	b.AppendNoTab(bs)
}

func (b *Builder) AppendNoTab(bs []byte) {
	b.buf = append(b.buf, bs...)
}

func (b *Builder) AppendLine() {
	b.buf = append(b.buf, '\n')
}

func (b *Builder) ApplyTabs() {
	for i := 0; i < b.numTab; i++ {
		b.buf = append(b.buf, '\t')
	}
}

//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func (b *Builder) copyCheck() {
	if b.addr == nil {
		b.addr = (*Builder)(noescape(unsafe.Pointer(b)))
	} else if b.addr != b {
		panic("lua-builder: illegal use of non-zero Builder copied by value")
	}
}
