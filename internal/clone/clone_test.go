package clone

import (
	"fmt"
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
		length, fileReadError := file.Read(fileBytes)

		if fileStatError != nil {
			t.Errorf("file.Read() error = %v, and we didn't ask for it", fileReadError)
			return
		}

		// TODO I could just somehow io.ReadAtLeast() https://gobyexample.com/reading-files

		// Q: What's the difference
		fmt.Printf("%d bytes: %s\n", length, string(fileBytes[:length]))
		fmt.Println(string(fileBytes))

		content := string(fileBytes)
		firstLine := stringutils.GetFirstLine(content)

		if firstLine != "Apache License" {
			t.Fatalf("Expected the first line of the LICENSE file to be 'Apache License', but it was '%s'.", firstLine)
		}
	})
}
