package config

import (
	"os"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
)

// loads configs
func init() {
	if os.Getenv("config") != "" {
		ConfigFile = os.Getenv("config")
	}

	LoadConfig(ConfigFile)
}

func LoadConfig(ConfigFile string) {
	if _, err := toml.DecodeFile(ConfigFile, &Global); err != nil {
		Fatal("Fail to load configs: " + err.Error())
	}

	Server = Global.Server
	Database = Global.Database
	Login = Global.Login
}

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		Views:       Templates(),
	}
}
