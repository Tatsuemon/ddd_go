package unit_of_work

import (
	"github.com/Tatsuemon/ddd_go/domain/repository"
	"github.com/jmoiron/sqlx"
)

type IunitOfWork interface {
	Commit() error
	Rollback() error
	GetUserRepository() repository.UserRepository
}

type UnitOfWork struct {
	Con         *sqlx.DB
	Transaction *sqlx.Tx
	// desposed       bool
	UserRepository repository.UserRepository
}

func NewUnitOfWork(con *sqlx.DB, repo repository.UserRepository) (IunitOfWork, error) {
	tx, err := con.Beginx()
	if err != nil {
		return nil, err
	}

	// TODO(Tatsuemon): repoがない時にnewする
	// if repo == nil {
	// 	repo =
	// }

	return &UnitOfWork{
		Con:         con,
		Transaction: tx,
		// desposed:       false,
		UserRepository: repo,
	}, nil
}

func (u UnitOfWork) Commit() error {
	err := u.Transaction.Commit()
	return err
}

func (u UnitOfWork) Rollback() error {
	err := u.Transaction.Rollback()
	return err
}

func (u UnitOfWork) GetUserRepository() repository.UserRepository {
	return u.UserRepository
}
