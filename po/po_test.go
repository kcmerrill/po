package po

import "testing"

func TestPo(t *testing.T) {
	p := NewPo()
	user := Attributes{
		"email":    "string",
		"username": "string",
		"visits":   "number",
	}
	p.AddEntity("kcmerrill", user)

	if e, err := p.Entity("kcmerrill"); err == nil {
		e.A("email").Get()
	} else {
		t.Fatalf(err.Error())
	}
}
