package configs

import (
	"os"

	"github.com/gofiber/template/html/v2"
)

// DefaultRoutes func for describe group of public routes.
func Templates() *html.Engine {

	engine := html.New("./views", ".html")
	//engine := jet.New("./views", ".jet.html")

	// Delims sets the action delimiters to the specified strings
	//engine.Delims("{{", "}}") // Optional. Default: engine delimiters

	// Reload the templates on each render, good for development

	if os.Getenv("DEBUG") == "true" {
		engine.Reload(true) // Optional. Default: false
		engine.Debug(true)
	}

	return engine
}
