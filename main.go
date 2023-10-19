package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/snxl/blockchain/domain/entity"
)

func main() {

	start := time.Now()

	wallet, err := entity.NewWallet()
	if err != nil {
		panic(err)
	}

	chain := entity.NewBlockchain()
	chain.Blocks = []*entity.Block{entity.NewBlock("Genesis", []byte{})}

	chain.AddBlock("first block after genesis", wallet.Address)
	chain.AddBlock("second block after genesis", wallet.Address)
	chain.AddBlock("third block after genesis", wallet.Address)

	fmt.Println()
	fmt.Printf("Wallet public key: %x\n", wallet.PublicKey)
	fmt.Printf("Wallet private key: %x\n", wallet.PrivateKey)
	fmt.Printf("Wallet address: %x\n", wallet.Address)
	fmt.Println()

	for _, block := range chain.Blocks {

		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
		for _, transaction := range block.Transaction {
			fmt.Printf("transaction: %v\n", transaction)
		}

		pow := entity.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

	elapsed := time.Since(start).Truncate(time.Second)
	fmt.Printf("\n\nend process in: %v\n\n", elapsed)
}
