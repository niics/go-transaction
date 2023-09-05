# Go-Transaction

## Install
```
go get github.com/niics/go-transaction
```

## Interface
```
type GoTransaction interface {
	BroadcastTransaction(endPoint string, reqBody CresteTxRequest) (*CreateTxResponse, error)
	GetTransactionStatus(endPoint string, req TxStatusResquest) (*TxStatusResponse, error)
}
```

## Data Type
```
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
```

## Usage
```
  // init package
  config := gotransaction.Init(gotransaction.Config{
    endPoint: "https://mock-node-wgqbnxruha-as.a.run.app",
	  privateKey: "MIIBOgIBAAJBAKj34GkxFhD90vcNLYLInFEX6Ppy1tPf9Cnzj4p4WGeKLs1Pt8QuKUpRKfFLfRYC9AIKjbJTWit+CqvjWYzvQwECAwEAAQJAIJLixBy2qpFoS4DSmoEmo3qGy0t6z09AIJtH+5OeRV1be+N4cDYJKffGzDa88vQENZiRm0GRq6a+HPGQMd2kTQIhAKMSvzIBnni7ot/OSie2TmJLY4SwTQAevXysE2RbFDYdAiEBCUEaRQnMnbp79mxDXDf6AU0cN/RPBjb9qSHDcWZHGzUCIG2Es59z8ugGrDY+pxLQnwfotadxd+Uyv/Ow5T0q5gIJAiEAyS4RaI9YG8EWx/2w0T67ZUVAw8eOMB6BIUg0Xcu+3okCIBOs/5OiPgoTdSy7bcF9IGpSE8ZgGKzgYQVZeN97YE00",
  })

  // braodcast transaction
  gotransaction.BroadcastTransaction(config.EndPoint, gotransaction.CresteTxRequest{
    Symbol: "ETH",
    Price: 4500,
    Timestamp: 1678912345,
  })

  // Check Transaction Status
  gotransaction.GetTransactionStatus(config.EndPoint, gotransaction.CresteTxRequest{
    TxHash: "e97ae379d666eed7576bf285c451baf9f0edba93ce718bf15b06c8a85d07b8d1",
  })
```

