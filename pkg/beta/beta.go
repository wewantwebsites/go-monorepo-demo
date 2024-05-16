package beta

import (
	"fmt"

	"github.com/wewantwebsites/pkg/alpha"
)

func Hello (name string) string {
	countMap := alpha.CoinCount(10)
	return fmt.Sprintf("Hello, %s %d", name, countMap[true])
}