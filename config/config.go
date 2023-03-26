package config

import (
	"fmt"
	"github.com/caarlos0/env/v7"
)

type Config struct {
	LogLevel   string     `evn:"LOG_LEVEL" envDefault:"DEBUG"`
	LdapConfig LdapConfig `envPrefix:"LDAP"`
}

type LdapConfig struct {
}

func ReadConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("error reading config %+v", err)
	}

	return &cfg, nil
}
