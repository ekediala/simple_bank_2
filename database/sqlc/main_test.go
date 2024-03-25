package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	conn, err := sql.Open(DatabaseConfig.DB_DRIVER, DatabaseConfig.DB_SOURCE)

	if err != nil {
		log.Fatalf("error conecting to database: %v", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
