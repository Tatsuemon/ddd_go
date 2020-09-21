package unit_of_work

import (
	"github.com/Tatsuemon/ddd_go/domain/repository"
	"github.com/jmoiron/sqlx"
)

type IunitOfWork interface {
	Commit() (error)
	Rollback() (error)
}


type UnitOfWork struct {
	 con *sqlx.DB
	 transaction *sqlx.Tx
	 despose bool
	 userRepository repository.UserRepository  
}

func NewUnitOfWork(con *sqlx.DB, repo repository.UserRepository) (IunitOfWork, error) {
	tx, err := con.Beginx()	
	if err != nil {
		return nil, err
	}
	return &UnitOfWork{
		con: con,
		transaction: tx,
		despose: false,
		userRepository: repo,
	}, nil
}

func (u UnitOfWork) Commit() error {
	err := u.transaction.Commit()
	return err
}

func (u UnitOfWork) Rollback() error {
	err := u.transaction.Rollback()
	return err
}

