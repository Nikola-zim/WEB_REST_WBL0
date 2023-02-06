package service

import (
	"WEB_REST_exm0302"
	"WEB_REST_exm0302/pkg/cash"
	"WEB_REST_exm0302/pkg/repository"
)

type Authorization interface {
	CreateUser(user WEB_REST_exm0302.User) (int, error)
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
	WriteInCash(inputJson WEB_REST_exm0302.Json) error
	ReadFromCash(id uint64) (WEB_REST_exm0302.Json, error)
	WriteInDB(inputJson WEB_REST_exm0302.Json) error
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
