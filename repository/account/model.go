package account

import "database/sql"

// Account Account
type Account struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type accountSQL struct {
	ID     sql.NullString
	Status sql.NullString
}

func (item *accountSQL) GetObject() Account {
	return Account{
		ID:     item.ID.String,
		Status: item.Status.String,
	}
}
