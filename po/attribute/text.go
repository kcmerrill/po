package attribute

import (
	"sync"
)

// Text contains a text value
type Text struct {
	lock  *sync.Mutex
	value string
}

// Get returns the text value
func (t *Text) Get() interface{} {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.value
}

// Set sets the text value
func (t *Text) Set(value string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.value = value
}

// Reset re-initalizes the text value
func (t *Text) Reset() {
	t.Set("")
}
