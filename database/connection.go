package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectSQL() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testing")
	if err != nil {
		fmt.Print(err.Error())
	}

	return db
}
