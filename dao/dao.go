package dao

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// NewMySQL new db and retry connection when has error.
func NewMySQL(dsn string) (db *sql.DB) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("open mysql error(%v)\n", err)
		panic(err)
	}
	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(5000)
	return
}
