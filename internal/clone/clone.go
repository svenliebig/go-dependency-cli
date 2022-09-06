package clone

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

// TODO add branch parameter
// TODO add FS return value
func GitClone(url string) {
	memory := memory.NewStorage()

	// TODO use this `nil` there, that should be a fileSystem
	repository, err := git.Clone(memory, nil, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		// TODO this should be a parameter
		// TODO decide if you want to go straight for the `main` branch or if
		//      you want to check if the branch exists. With a nice help.
		ReferenceName: "refs/heads/main",
	})

	treeObjects, _ := repository.TreeObjects()

	// TODO create an iterateOver func or someting, this annoys me
	//      https://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters
	for {
		// TODO What are TreeObjects and why are there so many of them
		//      The TreeObjects have the partly the same files.
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

		data, _ := ioutil.ReadAll(reader)
		reader.Close()

		str := string(data)
		fmt.Println(str) // <here
		fmt.Println(objectType.String())
		fmt.Println("Key:", key, "=>", "Element:", el)
	}

	log.Default().Println(config.Author.Email)
}
