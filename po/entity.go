package po

import (
	"fmt"
	"sync"

	"github.com/kcmerrill/po/po/attribute"
)

// Entity contains a list of attributes
type Entity struct {
	lock       *sync.Mutex
	Attributes map[string]attribute.Attribute
}

// NewEntity will create an empty Entity
func NewEntity() *Entity {
	return &Entity{
		lock:       &sync.Mutex{},
		Attributes: make(map[string]attribute.Attribute),
	}
}

// AddAttribute will add an attribute to an entity
func (e *Entity) AddAttribute(label, attributeType string) attribute.Attribute {
	a := attribute.New(attributeType)
	e.lock.Lock()
	e.Attributes[label] = a
	e.lock.Unlock()
	return a
}

// A will return an entity's attribute
func (e *Entity) A(label string) attribute.Attribute {
	if a, err := e.Attribute(label); err == nil {
		// sweet .. it exists
		return a
	}
	panic("Attribute '" + label + "' not found")
}

// Attribute will return an entity's attribute
func (e *Entity) Attribute(label string) (attribute.Attribute, error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	if a, exists := e.Attributes[label]; exists {
		return a, nil
	}

	return nil, fmt.Errorf("Unable to find attribute '" + label + "'")
}

// Export will return all attributes of an entity
func (e *Entity) Export() map[string]interface{} {
	a := make(map[string]interface{})
	e.lock.Lock()
	defer e.lock.Unlock()

	for label, attribute := range e.Attributes {
		a[label] = attribute.Get()
	}

	return a
}
