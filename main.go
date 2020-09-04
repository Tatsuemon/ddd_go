package main

import (
	"log"

	server "github.com/Tatsuemon/ddd_go/infrastructure"
	"github.com/Tatsuemon/ddd_go/infrastructure/datastore"
	"github.com/Tatsuemon/ddd_go/infrastructure/web/handler"
	"github.com/Tatsuemon/ddd_go/usecase"
)

func main() {
	var (
		datasource = "root:@/ddd_go"
		port       = 8080
	)

	db, err := datastore.NewDB(datasource)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	userRep := datastore.NewUserPersistence(db)
	userUC := usecase.NewUserUseCase(userRep)

	userHd := handler.NewUserHandler(userUC)

	s := server.NewServer()
	s.Init(db, userHd)
	s.Run(port)
}
