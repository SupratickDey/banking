package dto

type NewTransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	Amount          float64 `json:"amount"`
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
