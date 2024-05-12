package beta

import (
	"testing"
)
func TestHello(t *testing.T) {
	name := "Cas"
	expectedHello := "Hello, Cas"

	if got := Hello(name); got != expectedHello {
		t.Errorf("Hello(%s) = %v (expected: %s)", name, got, expectedHello)
	}
}