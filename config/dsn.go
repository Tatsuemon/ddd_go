package config

import (
	"fmt"
	"os"
)

func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),     // mysql
		os.Getenv("DB_PASSWORD"), // mysqlpass
		os.Getenv("DB_HOST"),     // db
		os.Getenv("DB_PORT"),     // 3306
		os.Getenv("DB_DATABASE"), // ddd_go
	) + "?parseTime=true&collation=utf8mb4_bin"
}
