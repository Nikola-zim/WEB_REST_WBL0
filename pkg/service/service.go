package service

import (
	"WEB_REST_exm0302/pkg/cash"
	"WEB_REST_exm0302/pkg/repository"
	"WEB_REST_exm0302/static"
)

type Authorization interface {
	CreateUser(user static.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type CashNumbers interface {
	AppendNumberInCash(newNum string) error
	ReadNumbersFromCash() (string, error)
}

type JsonRW interface {
	WriteInCash(inputJson static.Json) error
	ReadFromCash(id string) (static.Json, error)
	WriteInDB(inputJson static.Json) error
	WriteNatsJsonInCash([]byte) error
	WriteNatsJsonInDB([]byte) error
	RecoverCash() error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
	CashNumbers
	JsonRW
}

func NewService(repos *repository.Repository, testCash *cash.Cash) *Service {
	return &Service{
		//Передаём интерфейсы
		Authorization: NewAuthService(repos.Authorization),
		CashNumbers:   NewCashService(testCash.NumbersRW),
		JsonRW:        NewJsonService(testCash.CashJsonRW, repos.DBJsonRW),
	}
}
