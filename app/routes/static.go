package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Static(a *fiber.App) {
	// Create routes group.

	// Static files

	config := fiber.Static{
		//Compress:      true,
		ByteRange:     true,
		CacheDuration: 60 * time.Second,
		MaxAge:        3600,
	}

	a.Static("/static", "./public", config)
	a.Static("/favicon.ico", "./public/favicon.ico", config)
}
