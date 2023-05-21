package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/kelseyhightower/envconfig"
)

type Cfg struct {
	DataBaseURL string `envconfig:"DATABASE_URL" required:"true"`
}

func NewConfig() (Cfg, error) {
	cfg := Cfg{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		return Cfg{}, errors.New(err.Error())
	}

	return cfg, nil
}
