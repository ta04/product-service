// product-service/main.go

package main

import (
	"log"

	"github.com/SleepingNext/product-service/database"
	"github.com/SleepingNext/product-service/handler"
	productPB "github.com/SleepingNext/product-service/proto"
	"github.com/SleepingNext/product-service/repository/postgres"
	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
)

func main() {
	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.srv.product"),
	)

	// Initialize the service
	s.Init()

	// Connect to postgres
	db, err := database.OpenPostgresConnection()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create a new handler
	h := handler.NewHandler(&postgres.Repository{
		DB: db,
	})

	// Register the handler
	productPB.RegisterProductServiceHandler(s.Server(), h)

	// Run the service
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
