package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Database struct {
	DbConnection          string
	DbPoolMaxConnLifeTime string
	DbPoolMaxConns        int
}

type Config struct {
	Database Database
}

var Configuration Config

func setDefaultConfig() {
	viper.SetDefault("Database.DbPoolMaxConnLifeTime", "1h")
	viper.SetDefault("Database.DbPoolMaxConns", 4)
}

func InitConfig() {
	// --- defaults
	setDefaultConfig()

	viper.Unmarshal(&Configuration)

	dbconnSrc := "config file"

	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		Configuration.Database.DbConnection = dbURL
		dbconnSrc = "environment variable DATABASE_URL"
	}

	log.Infof("Using database connection info from %v", dbconnSrc)

}
