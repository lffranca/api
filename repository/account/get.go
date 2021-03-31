package account

import "github.com/lffranca/api/repository"

func Get() (AccountRepository, error) {
	db, errDB := repository.GetDB()
	if errDB != nil {
		return nil, errDB
	}

	return &sqliteRepository{db: db}, nil
}
