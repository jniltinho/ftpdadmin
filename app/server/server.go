package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jniltinho/ftpdadmin/app/config"
)

type FiberServer struct {
	*fiber.App
	//db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(config.FiberConfig()),
		//db:  database.New(),
	}

	return server
}
