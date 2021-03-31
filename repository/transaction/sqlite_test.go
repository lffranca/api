package transaction

import (
	"context"
	"log"
	"testing"

	"github.com/lffranca/api/environment"
	"github.com/lffranca/api/repository"
)

func TestTransacion(t *testing.T) {
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

	trans, errTrans := Get()
	if errTrans != nil {
		t.Error(errTrans)
	}

	transactions, errTransactions := trans.GetByAccountID(context.Background(), "3b1b1c0c-5c68-4352-b937-e3c68b6b1b16")
	if errTransactions != nil {
		t.Error(errTransactions)
	}

	log.Println(transactions)
}
