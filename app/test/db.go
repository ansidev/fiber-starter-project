package test

import (
	"time"

	"github.com/ansidev/fiber-starter-project/pkg/db"
	"github.com/ansidev/fiber-starter-project/pkg/log"
	ep "github.com/fergusstrange/embedded-postgres"
	"go.uber.org/zap/zapio"
)

func GetTestDbConfig() (db.SqlDbConfig, *ep.EmbeddedPostgres) {
	dbConfig := db.SqlDbConfig{
		DbDriver:   "postgres",
		DbHost:     "localhost",
		DbPort:     9876,
		DbName:     "test",
		DbUsername: "test",
		DbPassword: "test",
	}

	l := &zapio.Writer{Log: log.L()}

	testDb := ep.NewDatabase(ep.DefaultConfig().
		Username(dbConfig.DbUsername).
		Password(dbConfig.DbPassword).
		Database(dbConfig.DbName).
		Version(ep.V14).
		Port(uint32(dbConfig.DbPort)).
		StartTimeout(45 * time.Second).
		Logger(l))

	return dbConfig, testDb
}
