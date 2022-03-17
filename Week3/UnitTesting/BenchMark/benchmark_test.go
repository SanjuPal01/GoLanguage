// Command to run : go test -bench=. (it want regex - . will help us to run all or you can also pass benchmark name)
// Note:  it require some time

// go test -bench=BenchmarkConcat

// Github repo to learn: https://github.com/MarioCarrion/videos/tree/34601c62f804cceed75e92dfbc30e5a52a118b32/2021/12/07

package benchmark

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Concat("a", "b")
	}
}

// Comparing Which Writing operation is Fast
func BenchmarkFmtSprintF(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str += fmt.Sprintf("%s%s", str, "a")
	}
}

func BenchmarkBytesBUffer(b *testing.B) {
	var buf bytes.Buffer
	for n := 0; n < b.N; n++ {
		buf.WriteString("a")
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	var builder strings.Builder
	for n := 0; n < b.N; n++ {
		builder.WriteString("a")
	}
}
