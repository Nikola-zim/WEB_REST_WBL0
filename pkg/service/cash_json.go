package service

import (
	"WEB_REST_exm0302/pkg/cash"
	"WEB_REST_exm0302/pkg/repository"
	"WEB_REST_exm0302/static"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type CashJsonService struct {
	cash cash.CashJsonRW
	repo repository.DBJsonRW
}

// В конструкторе принимаем интерфейс кеша
func NewJsonService(cash cash.CashJsonRW, repo repository.DBJsonRW) *CashJsonService {
	return &CashJsonService{
		cash: cash,
		repo: repo,
	}
}

func (cs *CashJsonService) ReadFromCash(id string) (static.Json, error) {
	return cs.cash.ReadFromCash(id)
}

func (cs *CashJsonService) WriteInCash(inputJson static.Json) error {
	return cs.cash.WriteInCash(inputJson)
}
func (cs *CashJsonService) WriteInDB(inputJson static.Json) error {

	return cs.repo.WriteInDB(inputJson)
}

func (cs *CashJsonService) WriteNatsJsonInCash(inputJson []byte) error {
	var inputJsonStruct static.Json
	errInputJson := json.Unmarshal(inputJson, &inputJsonStruct)
	if errInputJson != nil {
		logrus.Println("Начало ошибки")
		logrus.Fatalf("Ошибка преобразования Json: %s", errInputJson.Error())
		return errInputJson
	}
	return cs.cash.WriteInCash(inputJsonStruct)
}

func (cs *CashJsonService) RecoverCash() error {
	//var bufMap := make(map[string]static.Json)
	//var err error
	bufMap, errRead := cs.repo.ReadAllFromDB()
	if errRead != nil {
		logrus.Fatalf("Ошибка преобразования Json: %s", errRead)
		return errRead
	}
	for _, jsonStruct := range bufMap {
		err := cs.cash.WriteInCash(jsonStruct)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cs *CashJsonService) WriteNatsJsonInDB(inputJson []byte) error {
	var inputJsonStruct static.Json
	errInputJson := json.Unmarshal(inputJson, &inputJsonStruct)
	if errInputJson != nil {
		logrus.Println("Начало ошибки")
		logrus.Fatalf("Ошибка преобразования Json: %s", errInputJson.Error())
		return errInputJson
	}
	return cs.repo.WriteInDB(inputJsonStruct)
}
