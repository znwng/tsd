package printer

import (
	"fmt"
	"strings"
)

func PrintIndented(depth int, format string, args ...any) {
	fmt.Printf("%s", strings.Repeat("|   ", depth))
	fmt.Printf(format, args...)
}
