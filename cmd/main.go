package main

import "github.com/svenliebig/go-dependency-cli/internal/clone"

func main() {
	clone.GitClone("https://github.com/halimath/mini-httpd.git")
}
