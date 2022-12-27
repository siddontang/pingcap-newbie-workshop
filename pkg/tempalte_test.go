package pkg

import "testing"

func TestAdd(t *testing.T) {
	if v := CmpThenAdd(1, 1); v != 2 {
		t.Errorf("expect %v, but got %v", 2, v)
	}

	// miss covering other conditions
}
