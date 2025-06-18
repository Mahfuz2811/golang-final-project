package repositories

import (
	"database/sql"
	"final-golang-project/models"
	"fmt"
)

type MySQLUserRepositoy struct {
	db *sql.DB
}

func NewMySQLUserRepositoy(db *sql.DB) *MySQLUserRepositoy {
	return &MySQLUserRepositoy{
		db: db,
	}
}

func (r *MySQLUserRepositoy) Create(user models.User) error {
	query := "INSERT INTO users (username, email, password, is_verified, verification_token) VALUES(?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.Username, user.Email, user.PasswordHash, user.IsVerified, user.VerificationToken)
	if err != nil {
		fmt.Printf("error during user creation: %s", err)
	}

	return err
}

func (r *MySQLUserRepositoy) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, email, password, is_verified, verification_token FROM users where email = ?"
	row := r.db.QueryRow(query, email)

	var user models.User
	error := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.IsVerified, &user.VerificationToken)
	if error != nil {
		if error == sql.ErrNoRows {
			return nil, nil // User not found
		}

		fmt.Printf("error to retrive user by email: %s", error)

		return nil, error
	}

	return &user, nil
}
