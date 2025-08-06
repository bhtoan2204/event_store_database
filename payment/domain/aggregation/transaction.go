package aggregation

type TransactionAggregate struct {
	ID              int64   `json:"id"`
	TransactionCode string  `json:"transaction_code"`
	AccountNo       string  `json:"account_no"`
	Type            string  `json:"type"`
	Amount          float64 `json:"amount"`
	Reference       string  `json:"reference,omitempty"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}
