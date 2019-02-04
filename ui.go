package swaggerui

import (
	"net/http"

	"github.com/gobuffalo/packr"
)

// Box returns the Swagger UI packr box
func Box() *packr.Box {
	return packr.NewBox("./swagger-ui/dist")
}

// HandleAt sets up a http handle at the given path
func HandleAt(path string) {
	http.Handle(path, http.FileServer(Box()))
}
