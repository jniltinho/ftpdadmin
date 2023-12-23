package app

import (
	"github.com/jniltinho/ftpdadmin/app/config"
	"github.com/jniltinho/ftpdadmin/app/models"
	"github.com/jniltinho/ftpdadmin/app/routes"
	"github.com/jniltinho/ftpdadmin/app/utils"

	"github.com/gofiber/fiber/v2"
)

func InitServer() {

	//users, _ := app.GetUsers()
	//utils.PrettyJson(users, false)

	//user, _ := app.GetUsersByID(2)
	//utils.PrettyJson(user, false)

	// Define Fiber config.
	cfg := config.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(cfg)

	routes.Default(app)
	routes.AppV1(app)
	routes.Static(app)

	PrintUsers()

	// Start server
	print("  Server running on port " + config.Server.Addr)
	app.Listen(config.Server.Addr)
}

func PrintUsers() {
	users := models.Users{}
	result, _ := users.GetUsers()
	utils.PrettyJson(result, false)
}

func PrintGroups() {
	groups := models.Groups{}
	result, _ := groups.GetGroups()
	utils.PrettyJson(result, false)
}
