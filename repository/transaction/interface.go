package transaction

import "context"

// TransactionRepository TransactionRepository
type TransactionRepository interface {
	GetByAccountID(context.Context, string) ([]Transaction, error)
}
