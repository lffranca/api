package repository

import (
	"log"
	"testing"

	"github.com/lffranca/api/environment"
)

func TestAccount(t *testing.T) {
	env := &environment.Environment{
		SQLitePath: "/home/lucas/go/src/github.com/lffranca/api/app.db",
	}

	conn, errConn := Get(&GetInput{
		Type: ConnectionSQLite,
		Env:  env,
	})
	if errConn != nil {
		log.Panicln(errConn)
	}

	defer conn.Close()

}
