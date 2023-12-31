package gotransaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Package interfaces
type GoTransaction interface {
	BroadcastTransaction(endPoint string, reqBody CresteTxRequest) (*CreateTxResponse, error)
	GetTransactionStatus(endPoint string, req TxStatusResquest) (*TxStatusResponse, error)
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

type TxStatusResquest struct {
	TxHash string
}

type TxStatusResponse struct {
	TxStatus string
}

func Init(c Config) (*config, error) {
	conf := c.config()

	return conf, nil
}

func BroadcastTransaction(endPoint string, reqBody CresteTxRequest) (*CreateTxResponse, error) {
	createTxURL := fmt.Sprintf("%s/braodcast", endPoint)
	u, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	r, err := http.Post(createTxURL, "application/json", bytes.NewReader(u))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	resBody, err := io.ReadAll(r.Body)

	var response CreateTxResponse
	err = json.Unmarshal(resBody, &response)
	return &response, nil
}

func GetTransactionStatus(endPoint string, req TxStatusResquest) (*TxStatusResponse, error) {
	createTxURL := fmt.Sprintf("%s/check/%s", endPoint, req.TxHash)
	r, err := http.Get(createTxURL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	resBody, err := io.ReadAll(r.Body)

	var response TxStatusResponse
	err = json.Unmarshal(resBody, &response)
	return &response, nil
}
