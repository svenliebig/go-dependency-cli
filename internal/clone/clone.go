package clone

import (
	"io/fs"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/svenliebig/go-dependency-cli/internal/converter"
)

// TODO add branch parameter
func GitClone(url string, branch string) (fs.FS, error) {
	memory := memory.NewStorage()
	billyFs := memfs.New()

	// TODO use this `nil` there, that should be a fileSystem
	// repository, err :=
	git.Clone(memory, billyFs, &git.CloneOptions{
		URL: url,
		// Progress: os.Stdout,
		// TODO this should be a parameter
		// TODO decide if you want to go straight for the `main` branch or if
		//      you want to check if the branch exists. With a nice help.
		ReferenceName: "refs/heads/main",
		// TODO evaluate for performance
		// SingleBranch: true,
	})

	// fmt.Println(billyFs.Root())

	converter := converter.CreateBillyFilesystemConverter(billyFs)
	newFs := converter.ToFS()

	// b, _ := repository.Branch("main")
	// ref, _ := repository.Reference(plumbing.ReferenceName(b.Remote))

	// treeObjects, _ := repository.TreeObjects()

	// TODO create an iterateOver func or someting, this annoys me
	//      https://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters
	// for {
	// 	// TODO What are TreeObjects and why are there so many of them
	// 	//      The TreeObjects have the partly the same files.
	// 	//      https://stackoverflow.com/questions/54139971/what-does-a-tree-mean-in-git
	// 	treeObject, erro := treeObjects.Next()

	// 	if treeObject == nil {
	// 		break
	// 	}

	// 	if erro != nil {
	// 		break
	// 	}

	// 	files := treeObject.Files()

	// 	for {
	// 		file, erroo := files.Next()

	// 		if file == nil {
	// 			break
	// 		}

	// 		if erroo != nil {
	// 			break
	// 		}

	// 		fmt.Println(file.Name)
	// 	}
	// }

	// branches, _ := repository.Branches()
	// // var reference plumbing.ReferenceName := "/ref"

	// for {
	// 	branch, ok := branches.Next()

	// 	if branch == nil {
	// 		break
	// 	}

	// 	if ok != nil {
	// 		break
	// 	}

	// 	// reference := branch.Name()
	// 	fmt.Println(branch.Name())
	// }

	// ref, err := r.Reference(reference, false)

	// fmt.Println(ref.Name())

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// config, configError := repository.Config()

	// if configError != nil {
	// 	log.Fatal(configError.Error())
	// }

	// head, err3 := r.Head()

	// index, _ := memory.IndexStorage.Index()

	// for key, el := range index.Entries {
	// 	fmt.Println("Key:", key, "=>", "Element:", el)
	// }

	// This contains partly filenames, partly something that looks like file/directory permissions
	// for key, el := range memory.ObjectStorage.Trees {
	// 	reader, _ := el.Reader()
	// 	objectType := el.Type()

	// 	data, _ := ioutil.ReadAll(reader)
	// 	reader.Close()

	// 	str := string(data)
	// 	fmt.Println(str)
	// 	fmt.Println(objectType.String())
	// 	fmt.Println(key.String())
	// 	fmt.Println("Key:", key, "=>", "Element:", el)
	// }

	// I get partly the content of files >here, but  I cannot get the filenames
	// for key, el := range memory.ObjectStorage.Objects {
	// 	reader, _ := el.Reader()
	// 	objectType := el.Type()

	// 	// https://gobyexample.com/reading-files
	// 	// Q: What's the difference between this
	// 	data, _ := ioutil.ReadAll(reader)
	// 	str := string(data)

	// 	// Q: and this:
	// 	// objectBytes := make([]byte, el.Size())
	// 	// length, _ := reader.Read(objectBytes)
	// 	// str := string(objectBytes[:length])

	// 	reader.Close()

	// 	fmt.Println(str) // <here
	// 	fmt.Println(objectType.String())
	// 	fmt.Println("Key:", key, "=>", "Element:", el)
	// }

	// log.Default().Println(config.Author.Email)

	return newFs, nil
}

// TODO test what is fast, map or array search
