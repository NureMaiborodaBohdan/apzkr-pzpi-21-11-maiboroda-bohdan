package repository

import (
	"apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/jmoiron/sqlx"
)

type AuthMysql struct {
	db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user AlcoSafe.User) (int, error) {
	query := "INSERT INTO User (Username, Email, Password, Role, Name, Surname, Patronymic, Sex) VALUES (?, ?, ?, 'User', ?, ?, ?, ?)"
	result, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.Name, user.Surname, user.Patronymic, user.Sex)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthMysql) CreateAdmin(user AlcoSafe.User) (int, error) {
	query := "INSERT INTO User (Username, Email, Password, Role, Name, Surname, Patronymic, Sex) VALUES (?, ?, ?, 'Admin', ?, ?, ?, ?)"
	result, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.Name, user.Surname, user.Patronymic, user.Sex)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthMysql) GetUser(username, password string) (AlcoSafe.User, error) {
	var user AlcoSafe.User
	query := "SELECT UserID FROM User WHERE username=? and password=?"
	err := r.db.Get(&user, query, username, password)
	return user, err
}
