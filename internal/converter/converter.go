package converter

import (
	"io/fs"
	"log"

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
	// TODO this feels so wrong
	file billy.File
	info fs.FileInfo
}

func (f MyFile) Stat() (fs.FileInfo, error) {
	return f.info, nil
}

func (f MyFile) Read(b []byte) (int, error) {
	return f.file.Read(b)
}

func (f MyFile) Close() error {
	return f.file.Close()
}

func (myFs *MyOwnFS) Open(name string) (fs.File, error) {
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
	repositoryFs.files = converter.convertDirectory(infos)

	return repositoryFs
}

// TODO implement defer for closing stremas https://go.dev/blog/defer-panic-and-recover
func (converter *BillyFilesystemConverter) convertDirectory(fileInfos []fs.FileInfo) map[string]fs.File {
	m := map[string]fs.File{}

	for index := range fileInfos {
		fileInfo := fileInfos[index]

		if fileInfo.IsDir() {
			// TODO
			// cantina band
			// merge map response
			// Q: how do I want to forward the prefix
		} else {
			billyFile, err := converter.from.Open(fileInfo.Name())

			if err != nil {
				log.Fatal(err)
			}

			defer billyFile.Close()

			// TODO make() bytes with the specific length, should improve performance. Right? RIGHT GUYS?
			// var bytes []byte
			// _, err = billyFile.Read(bytes)

			// if err != nil {
			// 	log.Fatal(err)
			// }

			// Q:
			// So I basically just wrap the billyFile here into another file
			// because I am not capable of copying the content of an in memory
			// file to another one.
			// basically.
			// .
			file := new(MyFile)
			file.file = billyFile
			file.info = fileInfo
			m[billyFile.Name()] = file

			// billyFile.Close()

			// No, Create does actually create a file in the filesystem
			// osFile, err := os.Create(billyFile.Name())
			// osFile := os.NewFile(0, billyFile.Name())

			// if err != nil {
			// 	log.Fatal(err)
			// }
			// osFile := os.NewFile(uintptr(length), billyFile.Name())

			// _, err = osFile.Write(bytes)

			// if err != nil {
			// 	log.Fatal(err)
			// }

			// osFile.Close()

		}
	}

	return m
}
