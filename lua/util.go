package lua

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var nameDupMap = sync.Map{}
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func randName(ty string) string {
	var newName string
	for {
		newName = ty + strconv.FormatUint(r.Uint64(), 16)
		_, loaded := nameDupMap.LoadOrStore(newName, struct{}{})
		if !loaded {
			break
		}
	}
	return newName
}

func withTab(b Builder, f func()) {
	b.SetNumTab(b.NumTab() + 1)
	f()
	b.SetNumTab(b.NumTab() - 1)
}
