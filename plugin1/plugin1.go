package plugin1

import (
	"fmt"
	"net/http"

	"github.com/jakecoffman/go-pluggable-example/core"
)

func init() {
	fmt.Println("Plugin1 init")
	http.HandleFunc("/plugin1", core.RenderTemplate(func(r *http.Request) ([]string, interface{}) {
		return []string{"template/plugin1.html"}, map[string]string{
			"title": "My Title",
			"h1":    "My H1",
		}
	}))
}
