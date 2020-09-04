package di

import (
	"log"

	"github.com/Tatsuemon/ddd_go/domain/repository"
	"github.com/Tatsuemon/ddd_go/infrastructure/datastore"
	"github.com/Tatsuemon/ddd_go/infrastructure/web/handler"
	"github.com/Tatsuemon/ddd_go/usecase"
	"github.com/jmoiron/sqlx"
)

type UserContainer struct {
	Repository repository.UserRepository
	Usecase    usecase.UserUseCase
	Handler    handler.UserHandler
}

func NewUserContainer(env string, db *sqlx.DB) *UserContainer {
	switch env {
	case "development":
		log.Print("development")
		user_repository := datastore.NewUserPersistence(db)
		user_usecase := usecase.NewUserUseCase(user_repository)
		user_handler := handler.NewUserHandler(user_usecase)
		return &UserContainer{
			Repository: user_repository,
			Usecase:    user_usecase,
			Handler:    user_handler,
		}
	case "test":
		log.Print("test")
		user_repository := datastore.NewUserPersistence(db)
		user_usecase := usecase.NewUserUseCase(user_repository)
		user_handler := handler.NewUserHandler(user_usecase)
		return &UserContainer{
			Repository: user_repository,
			Usecase:    user_usecase,
			Handler:    user_handler,
		}
	default:
		log.Print("default")
		user_repository := datastore.NewUserPersistence(db)
		user_usecase := usecase.NewUserUseCase(user_repository)
		user_handler := handler.NewUserHandler(user_usecase)
		return &UserContainer{
			Repository: user_repository,
			Usecase:    user_usecase,
			Handler:    user_handler,
		}
	}
}
