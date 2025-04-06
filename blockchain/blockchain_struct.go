package blockchain

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/duckcoding00/go-blockchain/constants"
)

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

func (bs *BlockchainStruct) AddBlock(b *Block) {
	txInBlock := make(map[string]bool)
	for _, tx := range b.Transactions {
		txInBlock[tx.TransactionHash] = true
	}

	var newTransactionPool []*Transaction
	for _, tx := range bs.TransactionPool {
		if !txInBlock[tx.TransactionHash] {
			newTransactionPool = append(newTransactionPool, tx)
		}
	}

	bs.TransactionPool = newTransactionPool

	bs.Blocks = append(bs.Blocks, b)
}

// implementing power of work algorithm
func (bc *BlockchainStruct) PoW(miner string) *Block {
	// get prevHash
	prevHash := bc.Blocks[len(bc.Blocks)-1].Hash()
	nonce := 0
	for {
		// create newBlock
		guessBlock := NewBlock(prevHash, int64(nonce))

		// copy the txn pool
		for _, tx := range bc.TransactionPool {
			newTx := NewTransaction(tx.From, tx.To, tx.Value, tx.Data)
			guessBlock.AddTxtoBlock(newTx)
		}

		// guess the hash
		guessHash := guessBlock.Hash()
		desiredHash := strings.Repeat("0", constants.Mining_difficult)
		solutionHash := guessHash[2 : 2+constants.Mining_difficult]

		if desiredHash == solutionHash {
			rewardTx := NewTransaction(constants.BlockChain_Address, miner, constants.Mining_Reward, []byte{})
			rewardTx.Status = constants.Success
			guessBlock.Transactions = append(guessBlock.Transactions, rewardTx)
			bc.AddBlock(guessBlock)
			log.Print(bc.ToJson(), "\n\n")
			prevHash = bc.Blocks[len(bc.Blocks)-1].Hash()
			nonce = 0
			continue
		}

		nonce++
	}
}
