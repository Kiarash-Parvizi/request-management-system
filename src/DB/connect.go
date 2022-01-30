package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("mysql", "kia:aPass@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected to MySql")
}
