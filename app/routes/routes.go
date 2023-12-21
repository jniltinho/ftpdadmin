package routes

import (
	"github.com/jniltinho/ftpdadmin/app/configs"
	"github.com/jniltinho/ftpdadmin/app/handlers"
	"github.com/jniltinho/ftpdadmin/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AppV1(a *fiber.App) {
	jwt := middlewares.NewAuthMiddleware(configs.Secret)
	api := a.Group("/api") // /api

	v1 := api.Group("/v1")                           // /api/v1
	v1.Post("/login", handlers.LoginApi)             // /api/v1/login
	v1.Get("/protected", jwt, handlers.ProtectedApi) // /api/v1/protected
}

// DefaultRoutes func for describe group of public routes.
func Default(a *fiber.App) {

	a.Get("/nilton", func(c *fiber.Ctx) error {
		c.Set("HX-Redirect", "/")
		return c.SendString("Pagina de login Nilton")
	})

	a.Get("/", handlers.CheckSession, func(c *fiber.Ctx) error {
		//return c.SendString("Pagina Home")
		return c.Render("home", fiber.Map{})
	})

	a.Get("/home", handlers.CheckSession, func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{})
	})

	a.Get("/users", handlers.CheckSession, func(c *fiber.Ctx) error {
		return c.Render("users", fiber.Map{})
	})

	a.Get("/groups", handlers.CheckSession, func(c *fiber.Ctx) error {
		return c.SendString("New Page Grupos")
	})

	a.Get("/add_group", handlers.CheckSession, func(c *fiber.Ctx) error {
		return c.SendString("New Page Adicionar Grupo")
	})

	a.Get("/add_user", handlers.CheckSession, func(c *fiber.Ctx) error {
		return c.SendString("New Page Adicionar Usuario")
	})

	a.Post("/login", handlers.Login)

	a.Get("/login", func(c *fiber.Ctx) error {
		//return c.SendString("Pagina de login")
		return c.Render("login", fiber.Map{})
	})

	a.All("/logout", handlers.CheckSession, handlers.Logout)
}
