package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

func dbConnect() *pgxpool.Pool {

	dbconfig := dbConfig()

	db, err := pgxpool.ConnectConfig(context.Background(), dbconfig)
	if err != nil {
		log.Fatal(err)
	}
	dbName := dbconfig.ConnConfig.Config.Database
	dbUser := dbconfig.ConnConfig.Config.User
	dbHost := dbconfig.ConnConfig.Config.Host
	log.Infof("Connected as %s to %s @ %s", dbUser, dbName, dbHost)
	return db
}

func dbConfig() *pgxpool.Config {
	dbconf := Configuration.Database.DbConnection
	// disallow blank config for safety
	if dbconf == "" {
		log.Fatal("Blank DbConnection is disallowed for security reasons")
	}

	dbconfig, err := pgxpool.ParseConfig(Configuration.Database.DbConnection)
	if err != nil {
		log.Fatal(err)
	}
	// Read and parse connection lifetime
	dbPoolMaxLifeTime, errt := time.ParseDuration(Configuration.Database.DbPoolMaxConnLifeTime)
	if errt != nil {
		log.Fatal(errt)
	}
	dbconfig.MaxConnLifetime = dbPoolMaxLifeTime

	// Read and parse max connections
	dbPoolMaxConns := Configuration.Database.DbPoolMaxConns
	if dbPoolMaxConns > 0 {
		dbconfig.MaxConns = int32(dbPoolMaxConns)
	}

	// Read current log level and use one less-fine level
	dbconfig.ConnConfig.Logger = logrusadapter.NewLogger(log.New())
	levelString, _ := (log.GetLevel() - 1).MarshalText()
	pgxLevel, _ := pgx.LogLevelFromString(string(levelString))
	dbconfig.ConnConfig.LogLevel = pgxLevel

	return dbconfig
}
