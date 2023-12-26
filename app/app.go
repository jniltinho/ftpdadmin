package app

import (
	"fmt"

	"github.com/jniltinho/ftpdadmin/app/config"
	"github.com/jniltinho/ftpdadmin/app/models"
	"github.com/jniltinho/ftpdadmin/app/server"
	"github.com/jniltinho/ftpdadmin/app/utils"
)

func InitServer() {

	PrintUsers()

	server := server.New()
	server.RegisterFiberRoutes()
	port := config.Server.Addr
	err := server.Listen(port)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}

func PrintUsers() {
	result, _ := models.ListUsers()
	utils.PrettyJson(result, false)
}

func PrintGroups() {
	groups := models.Groups{}
	result, _ := groups.GetGroups()
	utils.PrettyJson(result, false)
}
