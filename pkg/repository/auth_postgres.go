package repository

import (
	"WEB_REST_exm0302"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Структура имплемитирует наш инфтерфейс репозитория и работает с БД
type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user WEB_REST_exm0302.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO%s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password) //сам запрос по плейсхолдерам
	//Записываем id нового пользователя
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (WEB_REST_exm0302.User, error) {
	var user WEB_REST_exm0302.User
	query := fmt.Sprintf("SELECT id FROM%s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
