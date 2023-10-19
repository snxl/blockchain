package entity

type BlockChain struct {
	Blocks []*Block
}

func NewBlockchain() *BlockChain {
	return &BlockChain{}
}

func (chain *BlockChain) AddBlock(data string, minerAddress string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := NewBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)

	transaction := NewTransaction()
	transaction.Sender = "Blockchain"
	transaction.Recipient = minerAddress
	transaction.Amount = 50.0

	new.Transaction = append(new.Transaction, transaction)
}
