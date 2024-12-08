package config

import (
	"database/sql"
	"errors"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	DB   *sql.DB
	once sync.Once
)


func PgConnect() *sql.DB {

	if DB != nil {
		return DB
	}

	once.Do(func() {
		db, err := sql.Open("postgres", "user=myuser password=mypassword host=postgres dbname=mydatabase sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		if db.Ping() != nil {
			log.Fatal(errors.New("Cannot Ping to the database"))
		}

		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		DB = db
	})

	return DB
}
