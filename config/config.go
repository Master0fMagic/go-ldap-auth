package config

import (
	"fmt"
	"github.com/caarlos0/env/v7"
)

type Config struct {
	LogLevel   string     `env:"LOG_LEVEL" envDefault:"DEBUG"`
	LdapConfig LdapConfig `envPrefix:"LDAP_"`
}

type LdapConfig struct {
	Url             string `env:"URL"`
	BindUser        string `env:"BIND_USER"`
	Password        string `env:"PASSWORD"`
	BaseDn          string `env:"BASE_DN"`
	UserFilter      string `env:"USER_FILTER"`
	UserAuthPattern string `env:"USER_AUTH_PATTERN"`
}

func ReadConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("error reading config %+v", err)
	}

	return &cfg, nil
}
