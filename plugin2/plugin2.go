package plugin2

import (
	"fmt"

	_ "github.com/jakecoffman/app/core"
)

func init() {
	fmt.Println("Plugin2 init")
}

func main() {
	fmt.Println("Plugin2 main")
}
