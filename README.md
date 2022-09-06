# go-dependency-cli

## build

`go build`

## run

`go run main.go`

## ressources

* [project file layout](https://github.com/golang-standards/project-layout)
* [golang in vs code](https://github.com/golang/vscode-go/wiki/debugging)

## TODO

* time measuring, difference between memory etc.

```go
now := time.Now()
elapsed := time.Since(now)
fmt.Printf("time elapsed for order %d: %s\n", order, elapsed)
```