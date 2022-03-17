package benchmark

import (
	"strings"
)

func Concat(a, b string) string {
	// Inefficient Way
	// return fmt.Sprintf("%s%s", a, b)

	// Efficient Way
	builder := strings.Builder{}
	builder.WriteString("a")
	builder.WriteString("b")
	return builder.String()
}
