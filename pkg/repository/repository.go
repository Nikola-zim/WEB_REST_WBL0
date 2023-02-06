package repository

import (
	"WEB_REST_exm0302"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user WEB_REST_exm0302.User) (int, error)
	GetUser(username, password string) (WEB_REST_exm0302.User, error)
}

type DBJsonRW interface {
	WriteInDB(inputJson WEB_REST_exm0302.Json) error
	ReadFromDB()
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	DBJsonRW
	TodoList
	TodoItem
}

// Инициализируем репозиторий в конструкторе
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		DBJsonRW:      NewJsonPostgres(db),
	}
}
