package persistence

import (
	"github.com/Tatsuemon/ddd_go/domain/model"
	"github.com/Tatsuemon/ddd_go/domain/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserPersistence struct {
	Conn *sqlx.DB
}

func NewUserPersistence(conn *sqlx.DB) repository.UserRepository {
	return &UserPersistence{Conn: conn}
}

func (r *UserPersistence) FindByID(id string) (*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil, nil
}

func (r *UserPersistence) FindByName(name string) (*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil, nil
}

func (r *UserPersistence) Store(user *model.User) error {
	// TODO(Tatsuemon): 処理
	return nil
}

func (r *UserPersistence) Update(user *model.User) error {
	// TODO(Tatsuemon): 処理
	return nil
}

func (r *UserPersistence) Delete(user *model.User) error {
	// TODO(Tatsuemon): 処理
	return nil
}
