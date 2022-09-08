package converter

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/go-git/go-billy/v5"
)

type FilesystemConverter interface {
	ToFS() fs.FS
}

type BillyFilesystemConverter struct {
	from billy.Filesystem
}

func CreateBillyFilesystemConverter(bfs billy.Filesystem) FilesystemConverter {
	converter := new(BillyFilesystemConverter)
	converter.from = bfs
	return converter
}

type MyOwnFS struct {
	files map[string]fs.File
}

type MyFile struct {
	// TODO this feels so wrong... edit: still a bit
	// Q: when I create a pointer out of it, it get 8 Byte worse /op . . ?!
	info   fs.FileInfo
	reader *strings.Reader

	// TODO maybe implement isClosed boolean?
}

func (f MyFile) Stat() (fs.FileInfo, error) {
	return f.info, nil
}

func (f MyFile) Read(b []byte) (int, error) {
	return f.reader.Read(b)
}

func (f MyFile) Close() error {
	// Q: What actually TODO when closing a file?
	// Q: implement isClosed probably?
	return f.reader.UnreadByte()
}

func (myFs *MyOwnFS) Open(name string) (fs.File, error) {
	// Q: / TODO maybe check with startsWith, if not, add //,  to prevent things.
	file := myFs.files[name]

	if file == nil {
		return nil, &fs.PathError{Op: "Open", Path: name}
	}

	return file, nil
}

// Q: Maybe I want a static ConverterFactory.create(billy billy.FileSystem)
//    I don't like that I set the billy filesystem after i initialize the converter
//    this should feel more like a static function, but I like that I do not have to
//    pass the billy.Fs inifinetly into the child functions.
func (converter *BillyFilesystemConverter) ToFS() fs.FS {
	root := converter.from.Root()
	infos, err := converter.from.ReadDir(root)

	if err != nil {
		log.Fatal(err)
	}

	repositoryFs := new(MyOwnFS)
	// files := converter.getFiles(infos, root)

	// for name, length := range files {

	// }

	repositoryFs.files = converter.convertDirectory(infos, root)

	return repositoryFs
}

// TODO implement defer for closing stremas https://go.dev/blog/defer-panic-and-recover
func (converter *BillyFilesystemConverter) convertDirectory(fileInfos []fs.FileInfo, dir string) map[string]fs.File {
	filesMap := map[string]fs.File{}

	for index := range fileInfos {
		fileInfo := fileInfos[index]

		if fileInfo.IsDir() {
			files, err := converter.from.ReadDir(converter.from.Join(dir, fileInfo.Name()))

			if err != nil {
				log.Fatal(err)
			}

			nestedMap := converter.convertDirectory(files, converter.from.Join(dir, fileInfo.Name()))

			for key, file := range nestedMap {
				filesMap[filepath.Clean(key)] = file
			}
		} else {
			fileName := converter.from.Join(dir, fileInfo.Name())
			billyFile, err := converter.from.Open(fileName)

			if err != nil {
				fmt.Printf("Error while trying to read file %s.", fileName)
				log.Fatal(err)
			}

			defer billyFile.Close()

			var bytes []byte = make([]byte, fileInfo.Size())
			length, err := billyFile.Read(bytes)

			if err != nil {
				log.Fatal(err)
			}

			file := new(MyFile)

			reader := strings.NewReader(string(bytes[:length]))

			file.reader = reader
			file.info = fileInfo
			filesMap[filepath.Clean(billyFile.Name())] = file
		}
	}

	return filesMap
}

// func (converter *BillyFilesystemConverter) getFiles(fileInfos []fs.FileInfo, dir string) map[string]int64 {
// 	files := map[string]int64{}

// 	for index := range fileInfos {
// 		fileInfo := fileInfos[index]
// 		if fileInfo.IsDir() {
// 			// this works on the first level
// 			files, err := converter.from.ReadDir(fileInfo.Name())

// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			nestedMap := converter.getFiles(files, fileInfo.Name())

// 			for name, length := range nestedMap {
// 				// files[converter.from.Join(dir, name)]
// 				fmt.Printf("nested: %s // %s -> %d", fileInfo.Name(), name, length)
// 			}
// 		} else {
// 			files[fileInfo.Name()] = fileInfo.Size()
// 		}
// 	}

// 	return files
// }
