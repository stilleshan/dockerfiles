package ga

import "testing"

func TestDetect(t *testing.T) {
	err := Detect()
	if err != nil {
		t.Error(err)
	}
}
