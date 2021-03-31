package account

import (
	"context"
	"database/sql"
)

type sqliteRepository struct {
	db *sql.DB
}

// Get Get
func (item *sqliteRepository) Get(ctx context.Context) ([]Account, error) {
	query := `select id, status from accounts;`

	rows, errRows := item.db.QueryContext(ctx, query)
	if errRows != nil {
		return nil, errRows
	}

	defer rows.Close()

	items := []Account{}
	for rows.Next() {
		itemSQL := accountSQL{}
		if err := rows.Scan(
			&itemSQL.ID,
			&itemSQL.Status,
		); err != nil {
			return nil, err
		}

		items = append(items, itemSQL.GetObject())
	}

	return items, nil
}

// GetByID GetByID
func (item *sqliteRepository) GetByID(ctx context.Context, id string) (*Account, error) {
	query := `select id, status from accounts where id = ?;`

	itemSQL := accountSQL{}
	if err := item.db.QueryRowContext(ctx, query, id).Scan(
		&itemSQL.ID,
		&itemSQL.Status,
	); err != nil {
		return nil, nil
	}

	itemToReturn := itemSQL.GetObject()

	return &itemToReturn, nil
}
