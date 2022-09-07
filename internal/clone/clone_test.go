package clone

import (
	"testing"

	"github.com/svenliebig/go-dependency-cli/internal/utils/stringutils"
)

func TestGitClone(t *testing.T) {

	t.Run("should clone the repository 'https://github.com/halimath/mini-httpd.git' and transform it into a fs.FS", func(t *testing.T) {
		fs, err := GitClone("https://github.com/halimath/mini-httpd.git", "main")

		if err != nil {
			t.Errorf("GitClone() error = %v, and we didn't ask for it", err)
			return
		}

		file, fileOpenError := fs.Open("LICENSE")

		if fileOpenError != nil {
			t.Errorf("fs.Open() error = %v, and we didn't ask for it", fileOpenError)
			return
		}

		stat, fileStatError := file.Stat()

		if fileStatError != nil {
			t.Errorf("file.Stat() error = %v, and we didn't ask for it", fileStatError)
			return
		}

		// Q: What's the difference
		fileBytes := make([]byte, stat.Size())
		_, fileReadError := file.Read(fileBytes)

		if fileStatError != nil {
			t.Errorf("file.Read() error = %v, and we didn't ask for it", fileReadError)
			return
		}

		// TODO I could just somehow io.ReadAtLeast() https://gobyexample.com/reading-files

		content := string(fileBytes)
		firstLine := stringutils.GetFirstLine(content)

		if firstLine != "Apache License" {
			t.Fatalf("Expected the first line of the LICENSE file to be 'Apache License', but it was '%s'.", firstLine)
		}
	})
}

// go test -benchmem -run=^$ -bench ^BenchmarkGitClone$ github.com/svenliebig/go-dependency-cli/internal/clone

/*

with billy removed:
BenchmarkGitClone-12    	       1	1204926400 ns/op	  586136 B/op	    3369 allocs/op

*/

func BenchmarkGitClone(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GitClone("https://github.com/halimath/mini-httpd.git", "main")
	}
}
