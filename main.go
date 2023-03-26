package main

import (
	"github.com/Master0fMagic/go-ldap-auth/config"
	"github.com/Master0fMagic/go-ldap-auth/storage/ldap"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)

	ldapService, err := ldap.New(cfg.LdapConfig)
	if err != nil {
		log.Fatal(err)
	}

	defer ldapService.Close()
	log.Infof("LDAP connection established")
}
