package config

import (
	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	ListenPort  string `envconfig:"LISTEN_PORT" default:"37004"`
	Environment string `envconfig:"ENV_NAME" default:"local"`
}

type DBConfig struct {
	Writer WriterDB
	Reader ReaderDB
}

type DBBase struct {
	DBMS           string `envconfig:"DBMS" default:"mysql"`
	User           string `envconfig:"DB_USER" default:"root"`
	Pass           string `envconfig:"DB_PASSWORD" default:"mysql"`
	Port           string `envconfig:"DB_PORT" default:"mysql"`
	DBName         string `envconfig:"DB_NAME" default:"mysql"`
	MaxConnections int    `envconfig:"DB_MAX_CONNECTIONS" default:"10"`
}

type WriterDB struct {
	DBBase
	Host string `envconfig:"DB_HOST" default:"mysql"`
}

type ReaderDB struct {
	DBBase
	Host string `envconfig:"READER_DB_HOST" default:"mysql"`
}

// Config represents application configuration.
type Config struct {
	Server ServerConfig
	DB     DBConfig
}

// NewConfig get configuration struct.
func NewConfig() (*Config, error) {
	conf := &Config{}

	if err := envconfig.Process("", conf); err != nil {
		return nil, err
	}

	return conf, nil
}
