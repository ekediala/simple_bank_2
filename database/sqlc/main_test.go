package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("error loading environment variables: %v", err)
	}

	dbUrl, found := os.LookupEnv("DATABASE_URL")

	if !found {
		log.Fatal("Database url not found in environment variables")
	}

	var (
		DB_DRIVER = "postgres"
		DB_SOURCE = fmt.Sprintf("%v?sslmode=disable", dbUrl)
	)

	conn, err := sql.Open(DB_DRIVER, DB_SOURCE)

	if err != nil {
		log.Fatalf("error conecting to database: %v", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
