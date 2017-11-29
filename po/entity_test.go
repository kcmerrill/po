package po

import "testing"

func TestEntity(t *testing.T) {
	e := NewEntity()
	e.AddAttribute("email", "string")
	e.AddAttribute("visits", "number")

	// lets test some defaults
	if a, err := e.Attribute("email"); err == nil {
		if a.Get() != "" {
			t.Fatalf("String should be empty by default")
		}
	} else {
		t.Fatalf(err.Error())
	}

	// lets try updating the string
	aBefore, _ := e.Attribute("email")
	aBefore.Set("kcmerrill@gmail.com")
	aAfter, _ := e.Attribute("email")
	if aAfter.Get() != "kcmerrill@gmail.com" {
		t.Fatalf("Expected email to be set to 'kcmerrill@gmail.com")
	}

	// lets check our A shortcut
	if e.A("visits").Get() != 0 {
		t.Fatalf("Expected number attribute to return 0")
	}

	e.A("visits").Set("1000")
	if e.A("visits").Get() != 1000 {
		t.Fatalf("Expected get() to return 1000")
	}
}
