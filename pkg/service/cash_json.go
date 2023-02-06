package service

import (
	"WEB_REST_exm0302"
	"WEB_REST_exm0302/pkg/cash"
	"WEB_REST_exm0302/pkg/repository"
)

type CashJsonService struct {
	cash cash.CashJsonRW
	repo repository.DBJsonRW
}

func (cs *CashJsonService) ReadFromCash(id uint64) (WEB_REST_exm0302.Json, error) {
	return cs.cash.ReadFromCash(id)
}

// В конструкторе принимаем интерфейс кеша
func NewJsonService(cash cash.CashJsonRW, repo repository.DBJsonRW) *CashJsonService {
	return &CashJsonService{
		cash: cash,
		repo: repo,
	}
}

func (cs *CashJsonService) WriteInCash(inputJson WEB_REST_exm0302.Json) error {
	return cs.cash.WriteInCash(inputJson)
}
func (cs *CashJsonService) WriteInDB(inputJson WEB_REST_exm0302.Json) error {

	return cs.repo.WriteInDB(inputJson)
}
