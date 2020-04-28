/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

package main

import (
	"log"

	"github.com/SleepingNext/product-service/config"
	"github.com/SleepingNext/product-service/database"
	"github.com/SleepingNext/product-service/handler"
	productPB "github.com/SleepingNext/product-service/proto"
	"github.com/SleepingNext/product-service/repository/postgres"
	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	name := config.MicroServiceName()
	port := config.MicroServicePort()

	registry := consul.NewRegistry()

	s := micro.NewService(
		micro.Name(name),
		micro.Address(port),
		micro.Registry(registry),
	)
	s.Init()

	db, err := database.OpenPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := handler.NewHandler(&postgres.Postgres{
		DB: db,
	})
	productPB.RegisterProductServiceHandler(s.Server(), h)

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
