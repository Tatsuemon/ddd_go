package datastore

import (
	"log"


	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func TxHandler(db *sqlx.DB, f func(*sqlx.Tx) error) (err error) {
	tx, err := db.Beginx()	
	if err != nil {
		return errors.Wrap(err, "start transaction failed")
	}

	defer func() {
		if p := recover(); p != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				log.Fatalf("rollback failed: %v", rollBackErr)	
			}
			log.Print("Rollback operation")
			err = errors.Wrap(err, "transaction operation failed")
		} else if err != nil {
			err = errors.Wrap(err, "transaction operation failed")
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = f(tx)
	return err
}
