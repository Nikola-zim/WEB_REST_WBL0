package service

import (
	"WEB_REST_exm0302/pkg/cash"
)

type CashService struct {
	cash cash.NumbersRW
}

// В конструкторе принимаем интерфейс кеша
func NewCashService(cash cash.NumbersRW) *CashService {
	return &CashService{cash: cash}
}

// Передаем число в кеш
func (cs *CashService) AppendNumberInCash(newNum string) error {
	return cs.cash.AppendNumberInCash(newNum)
}

// Читаем тестовую мапу из кеша
func (cs *CashService) ReadNumbersFromCash() (string, error) {
	return cs.cash.ReadNumbersFromCash()
}
