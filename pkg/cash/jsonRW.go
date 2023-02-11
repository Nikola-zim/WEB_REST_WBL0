package cash

import (
	"WEB_REST_exm0302/static"
	"errors"
	"github.com/sirupsen/logrus"
	"sync"
)

type CashJson struct {
	JsonMap map[string]static.Json
	mux     sync.RWMutex
	wg      sync.WaitGroup
}

func NewCashJsonRW() *CashJson {
	return &CashJson{
		JsonMap: make(map[string]static.Json),
	}
}

// Запись
func (cj *CashJson) WriteInCash(inputJson static.Json) error {
	defer cj.wg.Done()
	cj.wg.Add(1)
	cj.mux.Lock()
	//Проверка наличия элемента
	_, err := cj.JsonMap[inputJson.OrderUid]
	if err {
		logrus.Println("Ошибка записи в кеш: значение уже существует")
		cj.mux.Unlock()
		return nil
	}
	cj.JsonMap[inputJson.OrderUid] = inputJson
	cj.mux.Unlock()
	return nil
}

// Read From cash
func (cj *CashJson) ReadFromCash(id string) (static.Json, error) {
	cj.mux.RLock()
	defer cj.mux.RUnlock()
	desiredJson, err := cj.JsonMap[id]
	if err != true {
		logrus.Println("Ошибка чтения из кеша: такого UID не существует")
		return desiredJson, errors.New("dummy")
	}
	return desiredJson, nil
}
