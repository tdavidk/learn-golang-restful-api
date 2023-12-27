package app

import (
	"database/sql"
	"tdavidk/learn-golang-restful-api/helper"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/learn_golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
