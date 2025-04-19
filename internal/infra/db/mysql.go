package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// NewMySqlConnector initializes and returns a new MySQL database connection using the
// provided configuration. It verifies the connection by pinging the database.
//
// Returns:
//   - *sql.DB: a pointer to the initialized database connection.
//   - error: an error if the connection could not be established or ping fails.
//
// Example:
//
//	db, err := NewMySqlConnector()
//	if err != nil {
//	    log.Fatal(err)
//	}
func NewMySqlConnector() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "your_user",
		Passwd: "your_passwr",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "foo",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
