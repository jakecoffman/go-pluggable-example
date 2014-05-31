package main

import (
	"net/http"

	_ "github.com/jakecoffman/app/core"
	_ "github.com/jakecoffman/app/plugin1"
	_ "github.com/jakecoffman/app/plugin2"
)

func init() {

}

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
