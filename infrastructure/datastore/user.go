package datastore

import (
	"github.com/Tatsuemon/ddd_go/domain/model"
	"github.com/Tatsuemon/ddd_go/domain/repository"
	"golang.org/x/xerrors"

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
	return nil, xerrors.New("error test")
}

func (r *userPersistence) FindByID(id string) (*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil, xerrors.New("error test")
}

func (r *userPersistence) FindByName(name string) (*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil, xerrors.New("error test")
}

func (r *userPersistence) Store(user *model.User) (*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil, xerrors.New("error test")
}

func (r *userPersistence) Update(user *model.User) (*model.User, error) {
	// TODO(Tatsuemon): 処理
	return nil, xerrors.New("error test")
}

func (r *userPersistence) Delete(user *model.User) error {
	// TODO(Tatsuemon): 処理
	return xerrors.New("error test")
}
