package stringutils

import "testing"

func TestGetFirstLine(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFirstLine(tt.args.str); got != tt.want {
				t.Errorf("GetFirstLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -benchmem -run=^$ -bench ^BenchmarkGitFirstLine* github.com/svenliebig/go-dependency-cli/internal/utils/stringutils

/*

 BenchmarkGitFirstLine-12                51992270                23.41 ns/op            0 B/op          0 allocs/op
 BenchmarkGitFirstLineCut-12             46153312                24.32 ns/op            0 B/op          0 allocs/op
 BenchmarkGitFirstLineSplit-12           11909475               102.8 ns/op            48 B/op          1 allocs/op
 BenchmarkGitFirstLineSplitN-12          31238659                39.15 ns/op           16 B/op          1 allocs/op

*/

func BenchmarkGitFirstLine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetFirstLine("                 asodasdaosk               \nasdddddasdasd\n")
	}
}

func BenchmarkGitFirstLineCut(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetFirstLineCut("                 asodasdaosk               \nasdddddasdasd\n")
	}
}

func BenchmarkGitFirstLineSplit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetFirstLineSplit("                 asodasdaosk               \nasdddddasdasd\n")
	}
}

func BenchmarkGitFirstLineSplitN(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetFirstLineSplitN("                 asodasdaosk               \nasdddddasdasd\n")
	}
}

// go test -benchmem -run=^$ -bench ^BenchmarkGetStringContent* github.com/svenliebig/go-dependency-cli/internal/utils/stringutils

/*
BenchmarkGetStringContent-12                            11035849               110.2 ns/op           640 B/op          1 allocs/op
BenchmarkGetStringContentWithLengthEvaluated-12         10979017               106.1 ns/op           640 B/op          1 allocs/op
BenchmarkGetStringContentWithLengthParameter-12         11032280               108.5 ns/op           640 B/op          1 allocs/op
*/

const STR = "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet."

var STRING_BYTES = []byte(STR)
var STRING_BYTES_LENGTH = len(STRING_BYTES)

func BenchmarkGetStringContent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetStringContent(STRING_BYTES)
	}
}

func BenchmarkGetStringContentWithLengthEvaluated(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetStringContentWithLengthEvaluated(STRING_BYTES)
	}
}

func BenchmarkGetStringContentWithLengthParameter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetStringContentWithLengthParameter(STRING_BYTES, STRING_BYTES_LENGTH)
	}
}
