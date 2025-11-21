package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func Database() (db *sql.DB, err error) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "my-secret-pw",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "mysqldb",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db, nil
}
