package entity

type BlockChain struct {
    Blocks []*Block
}

func NewBlockchain () *BlockChain {
    return &BlockChain{}
}

func (chain *BlockChain) AddBlock(data string) {
    prevBlock := chain.Blocks[len(chain.Blocks)-1]
    new := NewBlock(data, prevBlock.Hash)
    chain.Blocks = append(chain.Blocks, new)
}
