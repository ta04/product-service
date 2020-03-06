// product-service/database/database.go

package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "database"
	port     = 5432
	user     = "sleepingnext"
	password = "kevin99123"
	dbname   = "products"
)

// OpenPostgresConnection is to connect to postgres database
func OpenPostgresConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)

	return db, err
}
