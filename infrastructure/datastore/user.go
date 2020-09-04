package infrastructure

import (
	"github.com/Tatsuemon/ddd_go/domain/model"
	"github.com/Tatsuemon/ddd_go/domain/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// repository.UserRepositoryを満たす
type userPersistence struct {
	conn *sqlx.DB
}

func NewUserPersistence(conn *sqlx.DB) repository.UserRepository {
	return &userPersistence{conn: conn}
}

func (r *userPersistence) FindAll() ([]*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil, nil
}

func (r *userPersistence) FindByID(id string) (*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil, nil
}

func (r *userPersistence) FindByName(name string) (*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil, nil
}

func (r *userPersistence) Store(user *model.User) (*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil
}

func (r *userPersistence) Update(user *model.User) (*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil
}

func (r *userPersistence) Delete(user *model.User) error {
	// TODO(Tatsuemon): 処理
	return nil
}
