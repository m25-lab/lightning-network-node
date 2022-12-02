package config

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Timeout  int
}
