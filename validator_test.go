package simplvalidator

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	valuesToTest := []struct {
		value string
		valid bool
	}{
		{value: "", valid:true},
		{value: "foo", valid:false},
		{value: "bar", valid: false},
	}
	for _, s := range valuesToTest {
		isEmpty := IsEmpty(s.value)
		if !isEmpty == s.valid {
			t.Fatalf("Expected to be %t but got %t" , s.valid ,isEmpty)
		}

	}
}

