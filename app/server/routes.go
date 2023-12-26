package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jniltinho/ftpdadmin/app/config"
	"github.com/jniltinho/ftpdadmin/app/handlers"
	"github.com/jniltinho/ftpdadmin/app/middlewares"
	"github.com/jniltinho/ftpdadmin/app/models"
)

var Data = fiber.Map{}

func (s *FiberServer) RegisterFiberRoutes() {
	s.AppV1Routes()
	s.DefaultRoutes()
	s.StaticRoutes()
}

func (s *FiberServer) AppV1Routes() {
	jwt := middlewares.NewAuthMiddleware(config.Login.Secret)
	api := s.Group("/api") // /api

	v1 := api.Group("/v1")                           // /api/v1
	v1.Post("/login", handlers.LoginApi)             // /api/v1/login
	v1.Get("/protected", jwt, handlers.ProtectedApi) // /api/v1/protected
}

// DefaultRoutes func for describe group of public routes.
func (s *FiberServer) DefaultRoutes() {

	s.Get("/nilton", func(c *fiber.Ctx) error {
		c.Set("HX-Redirect", "/")
		return c.SendString("Pagina de login Nilton")
	})

	s.Get("/", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		return c.Render("home", Data)
	})

	s.Get("/home", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		return c.Render("home", Data)
	})

	s.Get("/users", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		users, _ := models.ListUsers()
		Data["Users"] = users
		return c.Render("users", Data)
	})

	s.Get("/add_user", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		return c.Render("add_user", Data)
	})

	s.Get("/groups", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		groups, _ := models.ListGroups()
		Data["Groups"] = groups
		return c.Render("groups", Data)
	})

	s.Get("/add_group", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		return c.Render("add_group", Data)
	})

	s.Post("/login", handlers.Login)

	s.Get("/login", func(c *fiber.Ctx) error {
		//return c.SendString("Pagina de login")
		return c.Render("login", fiber.Map{})
	})

	s.All("/logout", handlers.CheckSession, handlers.Logout)
}
