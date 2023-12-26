package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) StaticRoutes() {
	// Create routes group.

	// Static files

	config := fiber.Static{
		//Compress:      true,
		ByteRange:     true,
		CacheDuration: 60 * time.Second,
		MaxAge:        3600,
	}

	s.Static("/static", "./public", config)
	s.Static("/favicon.ico", "./public/favicon.ico", config)
}
