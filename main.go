package main

import (
	"log"
	"sync"
	"time"

	"github.com/duckcoding00/go-blockchain/blockchain"
	"github.com/duckcoding00/go-blockchain/constants"
)

func init() {
	log.SetPrefix(constants.BlockchainName + ": ")
}

func main() {
	var wg sync.WaitGroup

	genesisBlock := blockchain.NewBlock("0x0", 0)
	//transaction := blockchain.NewTransaction("0x1", "0x2", 100, []byte{})
	blockchain := blockchain.NewBlockchain(*genesisBlock)

	log.Println(blockchain.ToJson())
	log.Print("starting mining...", "\n\n")

	wg.Add(1)
	go func() {
		defer wg.Done()
		blockchain.PoW("bob")
	}()

	// time.Sleep(2 * time.Second)
	// blockchain.AddTXtoTXPool(*transaction)

	time.Sleep(3 * time.Second)

	wg.Wait()
}
