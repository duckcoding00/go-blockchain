package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/duckcoding00/go-blockchain/constants"
)

type Transaction struct {
	From            string `json:"from"`
	To              string `json:"to"`
	Value           uint64 `json:"value"`
	Data            []byte `json:"data"`
	Status          string `json:"status"`
	TransactionHash string `json:"transaction_hash"`
}

func NewTransaction(from, to string, value uint64, data []byte) *Transaction {
	t := new(Transaction)
	t.From = from
	t.To = to
	t.Value = value
	t.Data = data
	t.TransactionHash = t.Hash()
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

func (t Transaction) Hash() string {
	tc, _ := json.Marshal(t)
	sum := sha256.Sum256(tc)
	hex := hex.EncodeToString(sum[:32])
	formattedHex := constants.Hex_Prev + hex
	return formattedHex
}
