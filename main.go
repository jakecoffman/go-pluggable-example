package main

import (
	"net/http"

	_ "github.com/jakecoffman/go-pluggable-example/core"
	_ "github.com/jakecoffman/go-pluggable-example/plugin1"
	_ "github.com/jakecoffman/go-pluggable-example/plugin2"
)

func init() {

}

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
