package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var _db *sqlx.DB

// MustConnect connects to the database, and panics if this fails
func MustConnect() {
	var err error
    _db, err = sqlx.Connect("mysql", "marty:marty@tcp(127.0.0.1:3306)/mytodos?parseTime=true&loc=Local&multiStatements=false")
	if err != nil {
		panic(err)
	}
}
