package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var DbId *sql.DB

func init() {
	dbId_conn, err := connectTCPSocket("id")
	if err != nil {
		log.Fatalf("Database Connection Error: %v\n", err)
	}
	DbId = dbId_conn
}

func connectTCPSocket(dbName string) (*sql.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_tcp.go: %s environment variable not set.", k)
		}
		return v
	}

	var (
		dbHost = mustGetenv("DB_HOST")
		dbPort = mustGetenv("DB_PORT")
		dbUser = mustGetenv("DB_USER")
		dbPass = mustGetenv("DB_PASS")
	)

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s",
		dbHost, dbPort, dbUser, dbPass, dbName)

	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return dbPool, nil
}

func Close() {
	if DbId != nil {
		DbId.Close()
	}

	log.Printf("The database connection has been closed")
}
