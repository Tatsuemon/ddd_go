package main

import (
	"log"

	di "github.com/Tatsuemon/ddd_go/di/containers"
	server "github.com/Tatsuemon/ddd_go/infrastructure"
	"github.com/Tatsuemon/ddd_go/infrastructure/datastore"
)

func main() {
	var (
		datasource = "root:@/ddd_go"
		port       = 8080
	)

	// TODO(Tatsuemon): config.DSN()を使用
	db, err := datastore.NewMysqlDB(datasource)
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
