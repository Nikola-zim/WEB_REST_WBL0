package cash

import "WEB_REST_exm0302"

type NumbersRW interface {
	AppendNumberInCash(newNum string) error
	ReadNumbersFromCash() (string, error)
}

type CashJsonRW interface {
	WriteInCash(inputJson WEB_REST_exm0302.Json) error
	ReadFromCash(id uint64) (WEB_REST_exm0302.Json, error)
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
