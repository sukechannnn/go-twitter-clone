package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	var dsn string

	if os.Getenv("ENV") != "production" {
		dsn = "host=localhost port=15432 dbname=go-twitter-clone-dev user=postgres sslmode=disable TimeZone=UTC"
	} else if os.Getenv("ENV") == "production" {
		var (
			dbUser                 = os.Getenv("DB_USER")                  // e.g. 'my-db-user'
			dbPwd                  = os.Getenv("DB_PASS")                  // e.g. 'my-db-password'
			instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
			dbName                 = os.Getenv("DB_NAME")                  // e.g. 'my-database'
		)
		socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
		if !isSet {
			socketDir = "/cloudsql"
		}

		dsn = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
