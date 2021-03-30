package account

import "context"

// AccountRepository AccountRepository
type AccountRepository interface {
	Get(context.Context) ([]Account, error)
	GetByID(context.Context, string) (*Account, error)
}
