package repository

import (
	"database/sql"
	"os"

	"github.com/RohanSinghCode/allMySubscription-api/internal/database"
	_ "github.com/lib/pq"
)

type Connection struct {
	DB *database.Queries
}

func (db *Connection) ConnectToDb() (*sql.DB, error) {
	conn, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}

	db.DB = database.New(conn)

	return conn, nil
}
