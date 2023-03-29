package synctest

import "sync"

type Counter struct {
	sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Value() int {
	return c.value
}
func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}
