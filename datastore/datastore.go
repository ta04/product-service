package datastore

import "database/sql"

func ConnectPostgres() (*sql.DB, error) {
	connStr := "user=sleepingnext dbname=products sslmode=disable password=kevin99123"
	db, err := sql.Open("postgres", connStr)

	return db, err
}