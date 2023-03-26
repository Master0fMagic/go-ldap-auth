package main

import (
	"github.com/Master0fMagic/go-ldap-auth/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Info(cfg)
}
