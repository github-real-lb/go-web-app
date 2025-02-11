package db

import (
	"os"
	"testing"

	"github.com/github-real-lb/bookings-web-app/util/config"
)

const DBConfigFilename = "./../db.config.json"

var testStore DatabaseStore

func TestMain(m *testing.M) {
	dbConfig, err := config.LoadDBConfig(DBConfigFilename)
	if err != nil {
		os.Exit(1)
	}

	testStore, err = NewPostgresDBStore(dbConfig.TestDBConnectionString)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
