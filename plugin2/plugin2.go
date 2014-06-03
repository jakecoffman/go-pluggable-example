package plugin2

import (
	"fmt"

	_ "github.com/jakecoffman/go-pluggable-example/core"
)

func init() {
	fmt.Println("Plugin2 init")
}

func main() {
	fmt.Println("Plugin2 main")
}
