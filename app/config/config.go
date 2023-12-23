package config

// global configs
var (
	Global   GlobalConfig
	Server   ServerConfig
	Database DatabaseConfig
	Login    LoginConfig
)

// ConfigFile is the default config file
var ConfigFile = "./config.toml"

// Struct de configuração do servidor
type ServerConfig struct {
	Version            string `toml:"version"`
	Addr               string `toml:"addr"`
	Mode               string `toml:"mode"`
	StaticDir          string `toml:"static_dir"`
	ViewDir            string `toml:"view_dir"`
	UploadDir          string `toml:"upload_dir"`
	MaxMultipartMemory int    `toml:"max_multipart_memory"`
}

// Struct de configuração do banco de dados
type DatabaseConfig struct {
	Dialect      string `toml:"dialect"`
	DSN          string `toml:"datasource"`
	Dir          string `toml:"dir"`
	Table        string `toml:"table"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxOpenConns int    `toml:"max_open_conns"`
}

// Struct de configuração de login
type LoginConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Blowfish string `toml:"blowfish"`
	Secret   string `toml:"secret"`
}

// Struct de configuração global
type GlobalConfig struct {
	Server   ServerConfig   `toml:"server"`
	Database DatabaseConfig `toml:"database"`
	Login    LoginConfig    `toml:"login"`
}
