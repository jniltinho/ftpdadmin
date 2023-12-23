package config

import (
	"os"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
)

const Secret = "secret"

var GetConfig GlobalConfig

// ConfigFile is the default config file
var ConfigFile = "./config.toml"

type GlobalConfig struct {
	Server struct {
		Version            string `toml:"version"`
		Addr               string `toml:"addr"`
		Mode               string `toml:"mode"`
		StaticDir          string `toml:"static_dir"`
		ViewDir            string `toml:"view_dir"`
		UploadDir          string `toml:"upload_dir"`
		MaxMultipartMemory int    `toml:"max_multipart_memory"`
	} `toml:"server"`
	Database struct {
		Dialect      string `toml:"dialect"`
		DSN          string `toml:"datasource"`
		Dir          string `toml:"dir"`
		Table        string `toml:"table"`
		MaxIdleConns int    `toml:"max_idle_conns"`
		MaxOpenConns int    `toml:"max_open_conns"`
	} `toml:"database"`
	Login struct {
		Username string `toml:"username"`
		Password string `toml:"password"`
		Blowfish string `toml:"blowfish"`
	} `toml:"login"`
}

func LoadConfig(ConfigFile string) {
	var conf GlobalConfig
	if _, err := toml.DecodeFile(ConfigFile, &conf); err != nil {
		Fatal("Fail to load configs: " + err.Error())
	}

	GetConfig = conf
}

// loads configs
func init() {
	if os.Getenv("config") != "" {
		ConfigFile = os.Getenv("config")
	}

	LoadConfig(ConfigFile)
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
