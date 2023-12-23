package routes

import (
	"github.com/jniltinho/ftpdadmin/app/config"
	"github.com/jniltinho/ftpdadmin/app/handlers"
	"github.com/jniltinho/ftpdadmin/app/middlewares"
	"github.com/jniltinho/ftpdadmin/app/models"

	"github.com/gofiber/fiber/v2"
)

func AppV1(a *fiber.App) {
	jwt := middlewares.NewAuthMiddleware(config.Login.Secret)
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

		result, _ := models.ListUsers()
		return c.Render("users", fiber.Map{"User": result})
	})

	a.Get("/groups", handlers.CheckSession, func(c *fiber.Ctx) error {
		return c.Render("groups", fiber.Map{})
	})

	a.Get("/add_group", handlers.CheckSession, func(c *fiber.Ctx) error {
		return c.Render("add_group", fiber.Map{})
	})

	a.Get("/add_user", handlers.CheckSession, func(c *fiber.Ctx) error {
		return c.Render("add_user", fiber.Map{})
	})

	a.Post("/login", handlers.Login)

	a.Get("/login", func(c *fiber.Ctx) error {
		//return c.SendString("Pagina de login")
		return c.Render("login", fiber.Map{})
	})

	a.All("/logout", handlers.CheckSession, handlers.Logout)
}
