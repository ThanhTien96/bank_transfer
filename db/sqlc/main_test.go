package db

import (
	"database/sql"
	"log"
	"os"
	"simplebank/utils"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	cfg, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}
	testDB, err = sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
