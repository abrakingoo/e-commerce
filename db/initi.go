package db

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
	err error
)

func InitDB() error{
	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}
	
	return nil
}