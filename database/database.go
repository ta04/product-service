// product-service/database/database.go

package database

import "database/sql"

// OpenPostgresConnection is to connect to postgres database
func OpenPostgresConnection() (*sql.DB, error) {
	connStr := "user=sleepingnext dbname=products sslmode=disable password=kevin99123"
	db, err := sql.Open("postgres", connStr)

	return db, err
}