package config

import (
	"os"

	"github.com/gofiber/template/html/v2"
)

// DefaultRoutes func for describe group of public routes.
func Templates() *html.Engine {

	var conf = GetConfig

	engine := html.New("./views", ".html")
	//engine := jet.New("./views", ".jet.html")

	// Delims sets the action delimiters to the specified strings
	//engine.Delims("{{", "}}") // Optional. Default: engine delimiters

	// Reload the templates on each render, good for development

	if os.Getenv("DEBUG") == "true" || conf.Server.Mode == "debug" {
		engine.Reload(true) // Optional. Default: false
		engine.Debug(true)
		print("Debug mode")
	}

	return engine
}
