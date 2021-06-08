package ga

import (
	"net/url"
	"testing"
)

func TestConcatURLValues(t *testing.T) {
	v1 := make(url.Values)
	v2 := make(url.Values)

	v1.Set("a", "1")
	v1.Set("b", "2")
	v2.Set("a", "2")
	v2.Set("c", "3")

	concatURLValues(v1, v2)

	if v1.Encode() != "a=1&b=2&c=3" {
		t.Fail()
	}

	if v2.Encode() != "a=2&c=3" {
		t.Fail()
	}
}
