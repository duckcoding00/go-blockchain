package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/duckcoding00/go-blockchain/constants"
)

type Block struct {
	PrevHash     string         `json:"prevhash"`
	Timestamp    int64          `json:"timestamp"`
	Nonce        int64          `json:"nonce"`
	Transactions []*Transaction `json:"transactions"`
}

func NewBlock(prevHash string, nonce int64) *Block {
	b := new(Block)
	b.PrevHash = prevHash
	b.Timestamp = time.Now().UnixNano()
	b.Nonce = nonce
	b.Transactions = []*Transaction{}
	return b
}

func (b Block) ToJson() string {
	res, err := json.Marshal(b)

	if err != nil {
		return err.Error()
	}

	return string(res)
}

func (b Block) Hash() string {
	bc, _ := json.Marshal(b)
	sum := sha256.Sum256(bc)
	hex := hex.EncodeToString(sum[:32])
	formattedHex := constants.Hex_Prev + hex
	return formattedHex
}

func (b *Block) AddTxtoBlock(tx *Transaction) {
	// verify transaction
	if !tx.VerifyTx() {
		tx.Status = constants.Failed
	} else {
		tx.Status = constants.Success
	}
	b.Transactions = append(b.Transactions, tx)
}
