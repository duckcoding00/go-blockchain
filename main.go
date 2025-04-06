package main

import (
	"log"

	"github.com/duckcoding00/go-blockchain/constants"
	"github.com/duckcoding00/go-blockchain/wallet"
)

func init() {
	log.SetPrefix(constants.BlockchainName + ": ")
}

func main() {
	// var wg sync.WaitGroup

	// genesisBlock := blockchain.NewBlock("0x0", 0)
	// //transaction := blockchain.NewTransaction("0x1", "0x2", 100, []byte{})
	// blockchain := blockchain.NewBlockchain(*genesisBlock)

	// log.Println(blockchain.ToJson())
	// log.Print("starting mining...", "\n\n")

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	blockchain.PoW("bob")
	// }()

	// // time.Sleep(2 * time.Second)
	// // blockchain.AddTXtoTXPool(*transaction)

	// time.Sleep(3 * time.Second)

	// wg.Wait()

	w, _ := wallet.NewWallet()
	log.Println(w.GetPrivateKeyHex())
	log.Println(w.GetPublicKeyHex())
	log.Println(w.GetAddress())

	nw := wallet.NewWalletFromPrivate(w.GetPrivateKeyHex())
	log.Println(nw.GetPrivateKeyHex())
	log.Println(nw.GetPublicKeyHex())
	log.Println(nw.GetAddress())

	log.Println(w.GetPrivateKeyHex() == nw.GetPrivateKeyHex())
	log.Println(w.GetPublicKeyHex() == nw.GetPublicKeyHex())
	log.Println(w.GetAddress() == nw.GetAddress())

}
