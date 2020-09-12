package main

import (
	"log"

	"github.com/Tatsuemon/ddd_go/config"
	di "github.com/Tatsuemon/ddd_go/di/containers"
	server "github.com/Tatsuemon/ddd_go/infrastructure"
	"github.com/Tatsuemon/ddd_go/infrastructure/datastore"
)

func main() {
	var port = 8080

	db, err := datastore.NewMysqlDB(config.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	userContainer := di.NewUserContainer("test", db.DB)
	s := server.NewServer()
	s.Init(userContainer.Handler)
	s.Run(port)
}
