package gotransaction

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

func New() Transaction {
	return Transaction{}
}

func CreateTransactio() error {
	return nil
}

func GetTransactionStatus() error {
	return nil
}
