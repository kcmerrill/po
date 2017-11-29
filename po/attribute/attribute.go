package attribute

import "sync"

// Attribute methods needed for the various types
type Attribute interface {
	Set(value string)
	Get() interface{}
	Reset()
}

// New creates an attribute of given attribute type
func New(attribute string) Attribute {
	switch attribute {
	case "text":
		fallthrough
	case "string":
		return &Text{lock: &sync.Mutex{}}
	case "number":
		fallthrough
	case "int":
		return &Number{lock: &sync.Mutex{}}
	}
	return &Text{lock: &sync.Mutex{}, value: attribute}
}
