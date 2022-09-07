package clone

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

type FilesystemConverter interface {
	Convert() fs.FS
}

type BillyFilesystemConverter struct {
	from billy.Filesystem
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
func (converter *BillyFilesystemConverter) Convert() fs.FS {
	root := converter.from.Root()
	infos, err := converter.from.ReadDir(root)

	if err != nil {
		log.Fatal(err)
	}

	files := converter.convertDirectory(infos)

	repositoryFs := new(MyOwnFS)
	repositoryFs.files = files

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

			m[billyFile.Name()] = file
		}
	}

	return m
}

// TODO add branch parameter
// TODO add FS return value
func GitClone(url string, branch string) (fs.FS, error) {
	memory := memory.NewStorage()
	billyFs := memfs.New()

	// TODO use this `nil` there, that should be a fileSystem
	repository, err := git.Clone(memory, billyFs, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		// TODO this should be a parameter
		// TODO decide if you want to go straight for the `main` branch or if
		//      you want to check if the branch exists. With a nice help.
		ReferenceName: "refs/heads/main",
		// TODO evaluate for performance
		SingleBranch: true,
	})

	fmt.Println(billyFs.Root())

	converter := new(BillyFilesystemConverter)
	converter.from = billyFs
	newFs := converter.Convert()

	fmt.Println(newFs)

	// dirInfo, err := billyFs.ReadDir(billyFs.Root())

	// for index := range dirInfo {
	// 	nfo := dirInfo[index]

	// 	if nfo.IsDir() {
	// 		// cantina band
	// 	} else {
	// 		fmt.Println(nfo.Name())
	// 		fmt.Println(nfo.Size())

	// 	}
	// }

	// b, _ := repository.Branch("main")
	// ref, _ := repository.Reference(plumbing.ReferenceName(b.Remote))

	treeObjects, _ := repository.TreeObjects()

	// TODO create an iterateOver func or someting, this annoys me
	//      https://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters
	for {
		// TODO What are TreeObjects and why are there so many of them
		//      The TreeObjects have the partly the same files.
		//      https://stackoverflow.com/questions/54139971/what-does-a-tree-mean-in-git
		treeObject, erro := treeObjects.Next()

		if treeObject == nil {
			break
		}

		if erro != nil {
			break
		}

		files := treeObject.Files()

		for {
			file, erroo := files.Next()

			if file == nil {
				break
			}

			if erroo != nil {
				break
			}

			fmt.Println(file.Name)
		}
	}

	branches, _ := repository.Branches()
	// var reference plumbing.ReferenceName := "/ref"

	for {
		branch, ok := branches.Next()

		if branch == nil {
			break
		}

		if ok != nil {
			break
		}

		// reference := branch.Name()
		fmt.Println(branch.Name())
	}

	// ref, err := r.Reference(reference, false)

	// fmt.Println(ref.Name())

	if err != nil {
		log.Fatal(err.Error())
	}

	config, configError := repository.Config()

	if configError != nil {
		log.Fatal(configError.Error())
	}

	// head, err3 := r.Head()

	index, _ := memory.IndexStorage.Index()

	for key, el := range index.Entries {
		fmt.Println("Key:", key, "=>", "Element:", el)
	}

	// This contains partly filenames, partly something that looks like file/directory permissions
	for key, el := range memory.ObjectStorage.Trees {
		reader, _ := el.Reader()
		objectType := el.Type()

		data, _ := ioutil.ReadAll(reader)
		reader.Close()

		str := string(data)
		fmt.Println(str)
		fmt.Println(objectType.String())
		fmt.Println(key.String())
		fmt.Println("Key:", key, "=>", "Element:", el)
	}

	// I get partly the content of files >here, but  I cannot get the filenames
	for key, el := range memory.ObjectStorage.Objects {
		reader, _ := el.Reader()
		objectType := el.Type()

		// https://gobyexample.com/reading-files
		// Q: What's the difference between this
		data, _ := ioutil.ReadAll(reader)
		str := string(data)

		// Q: and this:
		// objectBytes := make([]byte, el.Size())
		// length, _ := reader.Read(objectBytes)
		// str := string(objectBytes[:length])

		reader.Close()

		fmt.Println(str) // <here
		fmt.Println(objectType.String())
		fmt.Println("Key:", key, "=>", "Element:", el)
	}

	log.Default().Println(config.Author.Email)

	return newFs, nil
}

// TODO test what is fast, map or array search
