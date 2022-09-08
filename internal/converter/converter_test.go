package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
)

func TestThings(t *testing.T) {
	file := os.NewFile(uintptr(2), "Hello_World")
	stat, err := file.Stat()

	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(stat)
}

// Q: I don't know how I want to name these tests
func TestCreateBillyFilesystemConverterWithASingleFile(t *testing.T) {
	t.Run("should create a fs.FS from the test billy fs and be able to Open 'Hello_World'", func(t *testing.T) {
		testFs := memfs.New()
		billyFile, err := testFs.Create("Hello_World")

		if err != nil {
			t.Fatalf(err.Error())
		}

		content := []byte("hello to the world")
		length, err := billyFile.Write(content)

		if err != nil {
			t.Fatalf(err.Error())
		}

		// TODO err
		converter := CreateBillyFilesystemConverter(testFs)

		if err != nil {
			t.Fatalf(err.Error())
		}

		// TODO err
		fs := converter.ToFS()

		file, err := fs.Open("Hello_World")

		if err != nil {
			t.Fatalf(err.Error())
		}

		stat, err := file.Stat()

		if err != nil {
			t.Fatalf(err.Error())
		}

		if stat.Size() != int64(length) {
			t.Fatalf("expected %d to equal %d in filesize, but it did not.", stat.Size(), length)
		}
	})
}

func TestCreateBillyFilesystemConverterWithADirectory(t *testing.T) {
	t.Run("should create a fs.FS from the test billy fs and be able to Open 'Hello_World'", func(t *testing.T) {
		testFs := memfs.New()
		testFs.MkdirAll("dir", 0777)
		billyFile, err := testFs.Create(testFs.Join("dir", "Hello_World"))

		if err != nil {
			t.Fatalf(err.Error())
		}

		content := []byte("hello to the world")
		length, err := billyFile.Write(content)

		if err != nil {
			t.Fatalf(err.Error())
		}

		// TODO err
		converter := CreateBillyFilesystemConverter(testFs)

		if err != nil {
			t.Fatalf(err.Error())
		}

		// TODO err
		fs := converter.ToFS()

		file, err := fs.Open(filepath.Join("dir", "Hello_World"))

		if err != nil {
			t.Fatalf(err.Error())
		}

		stat, err := file.Stat()

		if err != nil {
			t.Fatalf(err.Error())
		}

		if stat.Size() != int64(length) {
			t.Fatalf("expected %d to equal %d in filesize, but it did not.", stat.Size(), length)
		}
	})
}

var TEST_FS = memfs.New()
var TEST_FILE, _ = TEST_FS.Create("Hello_World")
var TEST_FILE_LENGTH, _ = TEST_FILE.Write([]byte("hello to the world"))

/*

BenchmarkCreateBillyFilesystemConverter-12    	  757297	      1555 ns/op	     728 B/op	      22 allocs/op

With DEFER billyFile.Close()
BenchmarkCreateBillyFilesystemConverter-12    	  697292	      1615 ns/op	     752 B/op	      23 allocs/op

With billyFile removedfrom the new file
BenchmarkCreateBillyFilesystemConverter-12    	  681592	      1777 ns/op	     824 B/op	      26 allocs/op
BenchmarkCreateBillyFilesystemConverter-12    	  632793	      1766 ns/op	     824 B/op	      26 allocs/op

After implementing directory handling
BenchmarkCreateBillyFilesystemConverter-12    	  510915	      2124 ns/op	     952 B/op	      33 allocs/op

*/

func BenchmarkCreateBillyFilesystemConverter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateBillyFilesystemConverter(TEST_FS).ToFS()
	}
}
