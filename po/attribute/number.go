package attribute

import (
	"strconv"
	"sync"
)

// Number contains a number value
type Number struct {
	lock  *sync.Mutex
	value int
}

// Get returns the number value
func (n *Number) Get() interface{} {
	n.lock.Lock()
	defer n.lock.Unlock()
	return n.value
}

// Set sets the number value
func (n *Number) Set(value string) {
	n.lock.Lock()
	defer n.lock.Unlock()
	if v, err := strconv.Atoi(value); err == nil {
		n.value = v
	}
}

// Reset re-initalizes the number value
func (n *Number) Reset() {
	n.Set("0")
}
