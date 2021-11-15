# go-lua

go-lua supports scripting Lua for Redis in golang.

---

### Variable

```lua
-- Lua
local a = nil
local b = "hello world"
local c = 2021
local d = true

a = b
b = "hi world"

local t = {}
t[1] = 1994
t["name"] = "tae"
```

```go
// Go
_ = lua.NewLua(func (l *lua.DefaultBuilder) {
    a := l.Local(lua.Nil())
    b := l.Local(lua.Str("hello world"))
    c := l.Local(lua.Num(2021))
    d := l.Local(lua.Bool(true))
    
    l.Assign(a, b)
    l.Assign(b, lua.Str("hi world"))
    
    t := l.Local(lua.Table())
    l.Assign(lua.At(t, lua.Num(1)), lua.Num(1994))
    l.Assign(lua.At(t, lua.Str("name")), lua.Str("tae"))
})
```

### Function Call

```lua
-- Lua
local a = "hello"
print(a, "world")
```

```go
// Go
_ = lua.NewLua(func (l *lua.DefaultBuilder) {
    a := l.Local(lua.Str("hello"))
    lualib.Print(a, lua.Str("world"))
})
```

### IF Statement

```lua
-- Lua
local a = 1
local b = 2
local c = 0
if a < b then
    c = a
else
    c = b
end
```

```go
// Go
_ = lua.NewLua(func (l *lua.DefaultBuilder) {
    a := l.Local(lua.Num(1))
    b := l.Local(lua.Num(2))
    c := l.Local(lua.Num(0))
    l.If(lua.Cond(a, "<", b)).Then(func() {
    	l.Assign(c, a)
    }).Else(func() {
        l.Assign(c, b)	
    }).End()
})
```

### For
```lua
-- Lua
for i = 1,10,1
do
    print(i)
end

for i = 10,1,-1
do
    if i >= 5 then
        print(i)
    end
end
```

```go
// Go
_ = lua.NewLua(func (l *lua.DefaultBuilder) {
    l.For(1, 10, 1).Do(func(i lua.Variable) {
    	lualib.Print(i)
    })
    
    a := l.Local(lua.Num(5))
    l.For(10, 1, -1).Do(func(i lua.Variable) {
    	l.If(lua.Cond(a, ">=", i)).Then(func() {
            lualib.Print(i)	
        }).End()
    })
})
```

### Return

```lua
-- Lua
return "bye world"
return "hello","world",100
```

```go
// Go
_ = lua.NewLua(func (l *lua.DefaultBuilder) {
    l.Return(lua.Str("bye world"))
    l.Return(lua.Str("hello"), lua.Str("world"), lua.Num(100))
})
```

### Any

```lua
-- Lua
local a = 5
print(a)
```

```go
// Go
_ = lua.NewLua(func (l *lua.DefaultBuilder) {
    l.Append([]byte("local a = 5"))
    l.AppendLine()
    l.Append([]byte("print(a)"))
})
```

---
## Features

- Basic Syntax
- ~~Operator~~ - planned
- Loop
- ~~Iterator~~ - planned
- Function Call
- ~~Function Definition In Lua~~ - planned
- ~~String Method~~ - planned
- Array
- Table
- ~~Module~~ - planned
- ~~Redis Lua Library~~ - planned