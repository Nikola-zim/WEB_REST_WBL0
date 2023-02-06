package cash

type NumbersRW interface {
	AppendNumberInCash(newNum string) error
	ReadNumbersFromCash() (string, error)
}

type Cash struct {
	NumbersRW
}

func NewCashTest() *Cash {
	return &Cash{
		NumbersRW: NewTestCash(),
	}
}
