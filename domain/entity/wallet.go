package entity

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

type Wallet struct {
	PublicKey  []byte
	PrivateKey []byte
	Address    string
}

func NewWallet() (*Wallet, error) {
	private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	wallet := &Wallet{PrivateKey: private.D.Bytes(), PublicKey: public}
	wallet.GenerateAddress()

	return wallet, nil
}

func (w *Wallet) GenerateAddress() {
	shaHash := sha256.Sum256(w.PublicKey)
	address := hex.EncodeToString(shaHash[:])
	w.Address = address
}
