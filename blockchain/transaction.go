package blockchain

import (
	"encoding/json"
)

type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Value  uint64 `json:"value"`
	Data   []byte `json:"data"`
	Status string `json:"status"`
}

func NewTransaction(from, to string, value uint64, data []byte) *Transaction {
	t := new(Transaction)
	t.From = from
	t.To = to
	t.Value = value
	t.Data = data
	return t
}

func (t Transaction) ToJson() string {
	res, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}

	return string(res)
}

func (t Transaction) VerifyTx() bool {
	if t.Value == 0 {
		return false
	}

	// TODO check the signature
	return true
}
