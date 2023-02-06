package repository

import (
	"WEB_REST_exm0302"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user WEB_REST_exm0302.User) (int, error)
	GetUser(username, password string) (WEB_REST_exm0302.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// Инициализируем репозиторий в конструкторе
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
