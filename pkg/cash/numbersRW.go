package cash

import (
	"bytes"
	"fmt"
)

type TestCash struct {
	Testmap map[int]string
}

func NewTestCash() *TestCash {
	return &TestCash{
		Testmap: make(map[int]string),
	}
}

func (tc *TestCash) AppendNumberInCash(newNum string) error {
	tc.Testmap[len(tc.Testmap)] = newNum
	return nil
}

func (tc *TestCash) ReadNumbersFromCash() (string, error) {

	mapInString := createKeyValuePairs(tc.Testmap)
	return mapInString, nil
}

func createKeyValuePairs(m map[int]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}
