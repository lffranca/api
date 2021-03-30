package account

func Get() (AccountRepository, error) {
	// rep, errRep := repository.

	return &sqliteRepository{}, nil
}
