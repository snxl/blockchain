package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/snxl/blockchain/domain/entity"
)

func main() {

    start:= time.Now()

    chain := entity.NewBlockchain()
    chain.Blocks = []*entity.Block{entity.NewBlock("Genesis", []byte{})}


    chain.AddBlock("first block after genesis")
    chain.AddBlock("second block after genesis")
    chain.AddBlock("third block after genesis")

    for _, block := range chain.Blocks {

        fmt.Printf("Previous hash: %x\n", block.PrevHash)
        fmt.Printf("data: %s\n", block.Data)
        fmt.Printf("hash: %x\n", block.Hash)

        pow := entity.NewProofOfWork(block)
        fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()
    }


    elapsed := time.Since(start).Truncate(time.Second)
    fmt.Printf("\n\nend process in: %v\n\n", elapsed)
}