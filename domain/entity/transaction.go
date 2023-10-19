package entity

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
	Fee       float64
	Signature []byte
	Nonce     int
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
