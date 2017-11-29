package po

import (
	"fmt"
	"sync"
)

// Attributes is a simple mapping(eye candy)
type Attributes map[string]string

// Po holds all of the entities
type Po struct {
	lock     *sync.Mutex
	token    string
	Entities map[string]*Entity
}

// NewPo creates a new Po initialized
func NewPo() *Po {
	return &Po{
		lock:     &sync.Mutex{},
		Entities: make(map[string]*Entity),
	}
}

// E Shorthand for creating/updating an entity
func (p *Po) E(label string, attributes Attributes) *Entity {
	if e, err := p.Entity(label); err == nil {
		return e
	}
	return p.AddEntity(label, attributes)
}

// Entity will return an entity if it exists
func (p *Po) Entity(label string) (*Entity, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if e, exists := p.Entities[label]; exists {
		return e, nil
	}
	return nil, fmt.Errorf("Unable to find entity '" + label + "'")
}

// AddEntity will add an Entity to Po
func (p *Po) AddEntity(label string, attributes Attributes) *Entity {
	e := NewEntity()

	for attributeLabel, attributeType := range attributes {
		e.AddAttribute(attributeLabel, attributeType)
	}

	p.lock.Lock()
	p.Entities[label] = e
	p.lock.Unlock()

	return e
}
