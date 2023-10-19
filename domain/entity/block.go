package entity

import (
	"bytes"
	"encoding/gob"

	"github.com/snxl/blockchain/domain/shared/helpers"
)

type Block struct {
	Hash        []byte
	Data        []byte
	PrevHash    []byte
	Transaction []*Transaction
	Nonce       int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash, Nonce: 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	helpers.HandleErrors(err)

	return res.Bytes()
}

func (b *Block) Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	helpers.HandleErrors(err)

	return &block
}
