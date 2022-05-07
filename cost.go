package comm

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Cost struct {
	PrintID int64
	TStart  int64
	TLast   int64
}

func (c *Cost) Show() {
	id := atomic.AddInt64(&c.PrintID, 1)
	now := time.Now().UnixNano()

	fmt.Printf("[%d][%d] cost:%d  last:%d\n", id, c.TStart/1000000, ((int64)(now)-c.TStart)/1000000,
		((int64)(now)-c.TLast)/1000000)
	c.TLast = (int64)(now)
}
func (c *Cost) ShowEx(key string) {
	id := atomic.AddInt64(&c.PrintID, 1)
	now := time.Now().UnixNano()

	fmt.Printf("[%d][%d][%s] cost:%d  last:%d\n", id, c.TStart/1000000, key, ((int64)(now)-c.TStart)/1000000,
		((int64)(now)-c.TLast)/1000000)
	c.TLast = (int64)(now)
}

func NewCoster() *Cost {
	now := time.Now().UnixNano()
	return &Cost{PrintID: 0, TStart: (int64)(now), TLast: (int64)(now)}
}
