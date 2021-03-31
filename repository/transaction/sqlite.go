package transaction

import (
	"context"
	"database/sql"
)

type sqliteRepository struct {
	db *sql.DB
}

// GetByAccountID GetByAccountID
func (item *sqliteRepository) GetByAccountID(ctx context.Context, accountID string) ([]Transaction, error) {
	query := `select id, id_account, description, value from transactions where id_account = ?;`

	rows, errRows := item.db.QueryContext(ctx, query, accountID)
	if errRows != nil {
		return nil, errRows
	}

	defer rows.Close()

	items := []Transaction{}
	for rows.Next() {
		itemSQL := transactionSQL{}
		if err := rows.Scan(
			&itemSQL.ID,
			&itemSQL.AccountID,
			&itemSQL.Description,
			&itemSQL.Value,
		); err != nil {
			return nil, err
		}

		items = append(items, itemSQL.GetObject())
	}

	return items, nil
}
