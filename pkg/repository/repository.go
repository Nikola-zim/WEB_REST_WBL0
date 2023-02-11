package repository

import (
	"WEB_REST_exm0302/static"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user static.User) (int, error)
	GetUser(username, password string) (static.User, error)
}

type DBJsonRW interface {
	WriteInDB(inputJson static.Json) error
	ReadAllFromDB() (map[string]static.Json, error)
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
