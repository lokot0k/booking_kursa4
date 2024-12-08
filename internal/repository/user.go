package repository

import (
	"database/sql"
	"meeting-room-booking/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	rows, err := r.db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetByID(id int) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(user domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Name, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(user domain.User) error {
	_, err := r.db.Exec("UPDATE users SET username = $1, password = $2 WHERE id = $3", user.Name, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
