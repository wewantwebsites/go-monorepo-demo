package beta

import (
	"fmt"
	"testing"

	"github.com/wewantwebsites/go-monorepo-demo/pkg/alpha"
)
func TestHello(t *testing.T) {
	name := "Cas"
	expectedHello := fmt.Sprintf("Hello, %s %d", name, alpha.CoinCount(10)[true])

	if got := Hello(name); got != expectedHello {
		t.Errorf("Hello(%s) = %v (expected: %s)", name, got, expectedHello)
	}
}