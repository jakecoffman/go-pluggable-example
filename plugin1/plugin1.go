package plugin1

import (
	"fmt"

	_ "github.com/jakecoffman/app/core"
)

func init() {
	fmt.Println("Plugin1 init")
}
