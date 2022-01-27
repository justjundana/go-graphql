package user

import (
	"database/sql"
	"log"

	_models "github.com/justjundana/go-graphql/models"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetUsers() ([]_models.User, error) {
	var users []_models.User
	rows, err := ur.db.Query(`SELECT id, name, email, password FROM users ORDER BY id ASC`)
	if err != nil {
		log.Fatalf("Error")
	}

	defer rows.Close()

	for rows.Next() {
		var user _models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Fatalf("Error")
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) GetUser(id int) (_models.User, error) {
	var user _models.User

	row := ur.db.QueryRow(`SELECT id, name, email FROM users WHERE id = ?`, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) CreateUser(user _models.User) (_models.User, error) {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

	statement, err := ur.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(user _models.User) (_models.User, error) {
	query := `UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?`

	statement, err := ur.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) DeleteUser(user _models.User) (_models.User, error) {
	query := `DELETE FROM users WHERE id = ?`

	statement, err := ur.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}
