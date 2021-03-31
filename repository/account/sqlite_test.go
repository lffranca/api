package account

import (
	"context"
	"log"
	"testing"

	"github.com/lffranca/api/environment"
	"github.com/lffranca/api/repository"
)

func TestAccount(t *testing.T) {
	env := &environment.Environment{
		SQLitePath: "/home/lucas/go/src/github.com/lffranca/api/app.db",
	}

	conn, errConn := repository.Get(&repository.GetInput{
		Type: repository.ConnectionSQLite,
		Env:  env,
	})
	if errConn != nil {
		log.Panicln(errConn)
	}

	defer conn.Close()

	acc, errAcc := Get()
	if errAcc != nil {
		log.Panicln(errAcc)
	}

	accounts, errAccounts := acc.Get(context.Background())
	if errAccounts != nil {
		log.Panicln(errAccounts)
	}

	accountsItem, errAccountsItem := acc.GetByID(context.Background(), "3b1b1c0c-5c68-4352-b937-e3c68b6b1b16")
	if errAccountsItem != nil {
		log.Panicln(errAccountsItem)
	}

	log.Println(accounts)
	log.Println(accountsItem)
}
