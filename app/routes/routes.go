package routes

import (
	"github.com/jniltinho/ftpdadmin/app/config"
	"github.com/jniltinho/ftpdadmin/app/handlers"
	"github.com/jniltinho/ftpdadmin/app/middlewares"
	"github.com/jniltinho/ftpdadmin/app/models"

	"github.com/gofiber/fiber/v2"
)

var Data = fiber.Map{}

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
		Data["Page"] = c.Route().Path
		return c.Render("home", Data)
	})

	a.Get("/home", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		return c.Render("home", Data)
	})

	a.Get("/users", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		users, _ := models.ListUsers()
		Data["Users"] = users
		return c.Render("users", Data)
	})

	a.Get("/add_user", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		return c.Render("add_user", Data)
	})

	a.Get("/groups", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		groups, _ := models.ListGroups()
		Data["Groups"] = groups
		return c.Render("groups", Data)
	})

	a.Get("/add_group", handlers.CheckSession, func(c *fiber.Ctx) error {
		Data["Page"] = c.Route().Path
		return c.Render("add_group", Data)
	})

	a.Post("/login", handlers.Login)

	a.Get("/login", func(c *fiber.Ctx) error {
		//return c.SendString("Pagina de login")
		return c.Render("login", fiber.Map{})
	})

	a.All("/logout", handlers.CheckSession, handlers.Logout)
}
