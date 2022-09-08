package clone

import (
	"io/fs"
	"log"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/svenliebig/go-dependency-cli/internal/converter"
	"github.com/svenliebig/go-dependency-cli/internal/utils/timer"
)

// TODO add branch parameter
func GitClone(url string, branch string) (fs.FS, error) {
	memory := memory.NewStorage()
	billyFs := memfs.New()

	timer.Start("git.Clone")
	// Q: / TODO we could us the repository for error handling, if the branch was not found
	_, err := git.Clone(memory, billyFs, &git.CloneOptions{
		URL: url,
		// Progress: os.Stdout,
		// TODO this should be a parameter
		// TODO decide if you want to go straight for the `main` branch or if
		//      you want to check if the branch exists. With a nice help.
		ReferenceName: "refs/heads/main",

		// TODO evaluate for performance
		// done, did nothing for performance....
		// .... Because the repository you tested has only one branch smartass
		SingleBranch: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	timer.Stop("git.Clone")

	timer.Start("converter.toFS")
	converter := converter.CreateBillyFilesystemConverter(billyFs)
	newFs := converter.ToFS()
	timer.Stop("converter.toFS")

	// b, _ := repository.Branch("main")
	// ref, _ := repository.Reference(plumbing.ReferenceName(b.Remote))

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

	return newFs, nil
}
