package main

import (
	"fmt"
	"log"

	"github.com/svenliebig/go-dependency-cli/internal/clone"
	"github.com/svenliebig/go-dependency-cli/internal/utils/timer"
)

func main() {
	timer.Start("clone.GitClone")
	fs, err := clone.GitClone("https://github.com/halimath/mini-httpd.git", "main")
	timer.Stop("clone.GitClone")

	timer.Start("check")

	if err != nil {
		log.Fatal(err)
	}

	file, err := fs.Open("LICENSE")

	if err != nil {
		log.Fatal(err)
	}

	x, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(x.Name())
	timer.Stop("check")

	timer.Print("git.Clone")
	timer.Print("converter.toFS")
	timer.Print("clone.GitClone")
	timer.Print("check")
}
