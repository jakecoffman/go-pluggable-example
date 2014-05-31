package main

import (
	"fmt"

	_ "github.com/jakecoffman/app/core"
	_ "github.com/jakecoffman/app/plugin1"
	_ "github.com/jakecoffman/app/plugin2"
)

func init() {
	fmt.Println("App Init")
}

func main() {
	fmt.Println("App main")
}
