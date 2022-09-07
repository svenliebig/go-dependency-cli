package main

import (
	"fmt"
	"log"

	"github.com/svenliebig/go-dependency-cli/internal/clone"
)

func main() {
	fs, err := clone.GitClone("https://github.com/halimath/mini-httpd.git", "main")

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
}
