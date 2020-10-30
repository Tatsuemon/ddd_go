package datastore

import (
	"context"
	"database/sql"

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
	users := make([]*model.User, 0)
	if err := r.conn.Select(&users, "Select id, name, email FROM users"); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userPersistence) FindByID(id string) (*model.User, error) {
	user := model.User{}
	if err := r.conn.Get(&user, "Select id, name, email FROM users WHERE id = ?", id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userPersistence) FindByName(name string) (*model.User, error) {
	user := model.User{}
	if err := r.conn.Get(&user, "Select id, name, email FROM users WHERE name = ?", name); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userPersistence) Store(ctx context.Context, user *model.User) (*model.User, error) {
	var tx interface{
		Prepare(query string) (*sql.Stmt, error)
	}
	tx, ok := GetTx(ctx)
	if !ok {
		tx = r.conn
	}
	stmt, err := tx.Prepare("INSERT INTO `users` (id, name, email) VALUES(?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	_, err = stmt.Exec(user.ID, user.Name, user.Email)
	// TODO(Tatsuemon): エラーハンドリングをもっと詳細に
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userPersistence) Update(ctx context.Context, user *model.User) (*model.User, error) {
	var tx interface{
		Prepare(query string) (*sql.Stmt, error)
	}
	tx, ok := GetTx(ctx)
	if !ok {
		tx = r.conn
	}
	stmt, err := tx.Prepare("UPDATE `users` SET name=?, email=? WHERE id=?")
	if err != nil {
		return nil, err
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	_, err = stmt.Exec(user.Name, user.Email, user.ID)
	// TODO(Tatsuemon): エラーハンドリングをもっと詳細に
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userPersistence) Delete(user *model.User) error {
	stmt, err := r.conn.Prepare("DELETE FROM `users` WHERE id=?")

	if err != nil {
		return err
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	_, err = stmt.Exec(user.ID)
	// TODO(Tatsuemon): エラーハンドリングをもっと詳細に
	if err != nil {
		return err
	}

	return nil
}
