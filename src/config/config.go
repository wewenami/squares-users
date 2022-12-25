package config

import (
	"os"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DBName string `env:"DB_NAME" envDefault:"users"`
	DSName string `env:"DS_NAME" envDefault:"postgres"`
	DBURL  string `env:"DB_URL" envDefault:"postgres://users:userspass@localhost:5433/users?sslmode=disble"`

	ServerConfiguration
}

type ServerConfiguration struct {
	GrpcListenPort string `env:":4000"`
}

func Parse(v any) error {
	defer os.Clearenv()
	return env.Parse(v)
}
