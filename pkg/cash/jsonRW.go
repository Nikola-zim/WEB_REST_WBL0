package cash

import (
	"WEB_REST_exm0302"
)

type CashJson struct {
	JsonMap map[uint64]WEB_REST_exm0302.Json
}

func NewCashJsonRW() *CashJson {
	return &CashJson{
		JsonMap: make(map[uint64]WEB_REST_exm0302.Json),
	}
}

// Запись
func (cj *CashJson) WriteInCash(inputJson WEB_REST_exm0302.Json) error {
	cj.JsonMap[uint64(len(cj.JsonMap))] = inputJson
	return nil
}

// Read From cash
func (cj *CashJson) ReadFromCash(id uint64) (WEB_REST_exm0302.Json, error) {
	//TODO безопасный вызов
	desiredJson := cj.JsonMap[id]
	return desiredJson, nil
}
