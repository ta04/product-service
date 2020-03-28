// product-service/database/database.go

package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "ec2-52-86-33-50.compute-1.amazonaws.com"
	port     = 5432
	user     = "bguaglboeqykqh"
	password = "04cb699639a52efc8710fdd3e44680fe6eb22fbf1d973e300011d14e7df712d9"
	dbname   = "d81epgnhgvfe3s"
)

// OpenPostgresConnection is to connect to postgres database
func OpenPostgresConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)

	return db, err
}
