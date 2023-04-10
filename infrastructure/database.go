package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var id_db *sql.DB

func init() {
	id_db_conn, err := connectTCPSocket("id")
	if err != nil {
		log.Fatalf("Database Connection Error: %v\n", err)
	}

	id_db = id_db_conn
}

func GetDB(db_name string) *sql.DB {
	switch db_name {
	case "id":
		return id_db
	}

	return nil
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
	if id_db != nil {
		id_db.Close()
	}
}
