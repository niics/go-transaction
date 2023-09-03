package gotransaction

// Package interfaces
type GoTransaction interface {
	CreateTxResponse()
	GetTransactionStatus()
}

type Transaction struct {
}

type CresteTxRequest struct {
	Symbol    string
	Price     uint64
	Timestamp uint64
}

type CreateTxResponse struct {
	TxHash string
}

type TxStatusResponse struct {
	TxStatus string
}

func Init(c Config) error {
	// conf := c.config()

	return nil
}

func CreateTransactio() error {
	return nil
}

func GetTransactionStatus() error {
	return nil
}
