package main

import (
	"log"

	"github.com/duckcoding00/go-blockchain/blockchain"
	"github.com/duckcoding00/go-blockchain/constants"
)

func init() {
	log.SetPrefix(constants.BlockchainName + ": ")
}

func main() {
	block := blockchain.NewBlock("0x0", 0)
	log.Println(block.ToJson())
	log.Println("Hash Block", block.Hash())

	transaction := blockchain.NewTransaction("0x1", "0x2", 14, []byte{})
	log.Println(transaction.ToJson())

	genBlock := block
	genBlock.Transactions = append(genBlock.Transactions, transaction)
	blockchain := blockchain.NewBlockchain(*genBlock)
	log.Println(blockchain.ToJson())
}
