package config

import (
	"blog/pkg/utils"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App App
	}
	App struct {
		Port      string `env:"APP_PORT" envDefault:"9090" validate:"required"`
		JwtSecret string `env:"JWT_SECRET" envDefault:"secret" validate:"required"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./.env", cfg)
	if err != nil {
		return nil, err
	}
	err = cleanenv.ReadConfig("./.env", cfg)
	if err != nil {
		fmt.Println("error loading config")
		return nil, err
	}
	err = utils.ValidateDto(cfg)
	if err != nil {
		fmt.Println("error validating config")
		return nil, err

	}
	fmt.Println("config loaded successfully")

	return cfg, nil
}
