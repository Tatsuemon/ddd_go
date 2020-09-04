package datastore

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewDB(datasource string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", datasource)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}

	// 何かわからん
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	// 接続確認？
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	return db, nil
}
