package main

import (
	"database/sql"
	"log"
	"meeting-room-booking/internal/config"
	"reflect"
)

type migration interface {
	Up(db *sql.DB) error
	Down(db *sql.DB) error
}

func logMigrationSuccess(m migration) {
	log.Printf("[UP] %s", reflect.ValueOf(m).Type())
}

func main() {
	db := config.PgConnect()
	defer db.Close()
	log.Println("Starting database migration...")

	migrations := []migration{
		create_table_users{},
		create_table_bookings{},
	}
	if err := migrate(db, migrations); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully!")
}

func migrate(db *sql.DB, migrations []migration) error {
	for _, m := range migrations {
		if err := m.Up(db); err != nil {
			return err
		}

		logMigrationSuccess(m)
	}

	return nil
}

type create_table_users struct{}

func (c create_table_users) Up(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);`
	_, err := db.Exec(query)

	return err
}

func (c create_table_users) Down(db *sql.DB) error {
	return nil
}

type create_table_bookings struct{}

func (c create_table_bookings) Up(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS bookings (
		id SERIAL PRIMARY KEY,
		room_name VARCHAR(255) NOT NULL,
		start_time TIMESTAMP NOT NULL,
		end_time TIMESTAMP NOT NULL,
		user_id INTEGER REFERENCES users (id)
	);`
	_, err := db.Exec(query)

	return err
}

func (c create_table_bookings) Down(db *sql.DB) error {
	return nil
}
