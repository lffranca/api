package transaction

import "database/sql"

// Transaction Transaction
type Transaction struct {
	ID          string  `json:"id"`
	AccountID   string  `json:"id_conta"`
	Description string  `json:"descricao"`
	Value       float64 `json:"valor"`
}

type transactionSQL struct {
	ID          sql.NullString
	AccountID   sql.NullString
	Description sql.NullString
	Value       sql.NullFloat64
}

func (item *transactionSQL) GetObject() Transaction {
	return Transaction{
		ID:          item.ID.String,
		AccountID:   item.AccountID.String,
		Description: item.Description.String,
		Value:       item.Value.Float64,
	}
}
