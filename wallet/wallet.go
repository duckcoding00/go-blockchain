package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey `json:"private_key"`
	PublicKey  *ecdsa.PublicKey  `json:"public_key"`
}

func NewWallet() (*Wallet, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	wallet := new(Wallet)
	wallet.PrivateKey = privateKey
	wallet.PublicKey = &privateKey.PublicKey

	return wallet, nil
}

func NewWalletFromPrivate(privateKey string) *Wallet {
	pK := privateKey[2:]
	d := new(big.Int)
	d.SetString(pK, 16)

	// create public key
	var npk ecdsa.PrivateKey
	npk.D = d
	npk.PublicKey.Curve = elliptic.P256()
	npk.PublicKey.X, npk.PublicKey.Y = npk.PublicKey.Curve.ScalarBaseMult(d.Bytes())

	wallet := new(Wallet)
	wallet.PrivateKey = &npk
	wallet.PublicKey = &npk.PublicKey

	return wallet

}

func (w *Wallet) GetPrivateKeyHex() string {
	return fmt.Sprintf("0x%x", w.PrivateKey.D)
}

func (w *Wallet) GetPublicKeyHex() string {
	return fmt.Sprintf("0x%x%x", w.PublicKey.X, w.PublicKey.Y)
}

func (w *Wallet) GetAddress() string {
	hash := sha256.Sum256([]byte(w.GetPublicKeyHex()[2:]))
	hex := fmt.Sprintf("%x", hash[:])
	address := "0x" + hex[len(hex)-40:]

	return address
}
