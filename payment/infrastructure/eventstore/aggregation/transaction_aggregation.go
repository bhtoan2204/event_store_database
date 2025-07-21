package aggregation

type TransactionAggregation struct {
	ID        string `json:"id"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}
