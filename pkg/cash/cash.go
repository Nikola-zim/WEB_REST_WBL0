package cash

import (
	"WEB_REST_exm0302/static"
)

type NumbersRW interface {
	AppendNumberInCash(newNum string) error
	ReadNumbersFromCash() (string, error)
}

type CashJsonRW interface {
	WriteInCash(inputJson static.Json) error
	ReadFromCash(id string) (static.Json, error)
}

type Cash struct {
	NumbersRW
	CashJsonRW
}

func NewCashTest() *Cash {
	return &Cash{
		NumbersRW:  NewTestCash(),
		CashJsonRW: NewCashJsonRW(),
	}
}
