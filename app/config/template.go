package config

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/template/html/v2"
)

// DefaultRoutes func for describe group of public routes.
func Templates() *html.Engine {

	engine := html.New("./views", ".html")

	m := map[string]interface{}{
		"formatTime":   formatTime,
		"formatDate":   formatDate,
		"bytesToMB":    bytesToMB,
		"userDisabled": userDisabled,
		"expiryDate":   expiryDate,
	}
	engine.AddFuncMap(m)

	// Delims sets the action delimiters to the specified strings
	//engine.Delims("{{", "}}") // Optional. Default: engine delimiters

	// Reload the templates on each render, good for development

	if os.Getenv("DEBUG") == "true" || Server.Mode == "debug" {
		engine.Reload(true) // Optional. Default: false
		engine.Debug(true)
		print("Debug mode")
	}

	return engine
}

func formatTime(t time.Time) string {
	return t.Format("02-01-2006 15:04")
}

func formatDate(t time.Time, format string) string {
	return t.Format(format)
}

func bytesToMB(b uint64) string {
	convert := float64(b) / 1048576
	s := fmt.Sprintf("%2.1f", convert)
	return s

}

func userDisabled(s uint16) string {
	if s != 0 {
		return "Yes"
	}
	return "No"
}

func expiryDate() string {
	t := time.Now().Add(time.Hour * 24 * 365)
	return t.Format("2006-01-02 15:04:05")
}
