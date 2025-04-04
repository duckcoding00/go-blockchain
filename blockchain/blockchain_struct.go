package blockchain

import "encoding/json"

type BlockchainStruct struct {
	TransactionPool []*Transaction `json:"transaction_pool"`
	Blocks          []*Block       `json:"blocks"`
}

func NewBlockchain(genBlock Block) *BlockchainStruct {
	bs := new(BlockchainStruct)
	bs.TransactionPool = []*Transaction{}
	bs.Blocks = []*Block{}
	bs.Blocks = append(bs.Blocks, &genBlock)

	return bs
}

func (bs BlockchainStruct) ToJson() string {
	res, err := json.Marshal(bs)
	if err != nil {
		return err.Error()
	}

	return string(res)
}

func (bs *BlockchainStruct) AddTXtoTXPool(tx Transaction) {
	bs.TransactionPool = append(bs.TransactionPool, &tx)
}
