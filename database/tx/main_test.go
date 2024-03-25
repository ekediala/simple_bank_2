package tx_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	database "github.com/ekediala/simple_bank_2/database/sqlc"
	"github.com/ekediala/simple_bank_2/database/tx"
)

var testTx *tx.TxManager

func TestMain(m *testing.M) {

	conn, err := sql.Open(database.DatabaseConfig.DB_DRIVER, database.DatabaseConfig.DB_SOURCE)

	if err != nil {
		log.Fatalf("error conecting to database: %v", err)
	}

	testTx = tx.New(conn)

	os.Exit(m.Run())
}
