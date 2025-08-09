package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env      string `mapstructure:"env"`
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver   string         `mapstructure:"driver"`
	SQLite   SQLiteConfig   `mapstructure:"sqlite"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type SQLiteConfig struct {
	Path string `mapstructure:"path"`
}

type PostgresConfig struct {
	URL string `mapstructure:"url"`
}

type JWTConfig struct {
	Secret  string `mapstructure:"secret"`
	Expires string `mapstructure:"expires"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	// Add environment variable bindings
	viper.BindEnv("env", "ENV")
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("database.driver", "DATABASE_DRIVER")
	viper.BindEnv("database.sqlite.path", "DATABASE_SQLITE_PATH")
	viper.BindEnv("database.postgres.url", "DATABASE_POSTGRES_URL")
	viper.BindEnv("jwt.secret", "JWT_SECRET")
	viper.BindEnv("jwt.expires", "JWT_EXPIRES")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
