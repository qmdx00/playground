package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitMySQL() *sqlx.DB {
	dsn := "root@tcp(127.0.0.1:3306)/playground?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("connect to MySQL failed")
	}
	return db
}
