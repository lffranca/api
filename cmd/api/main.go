package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lffranca/api/api"
	"github.com/lffranca/api/environment"
	"github.com/lffranca/api/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARNING] Not use .env file")
	}

	env, errENV := environment.Get()
	if errENV != nil {
		log.Panicln(errENV)
	}

	conn, errConn := repository.Get(&repository.GetInput{
		Type: repository.ConnectionSQLite,
		Env:  env,
	})
	if errConn != nil {
		log.Panicln(errConn)
	}

	defer conn.Close()

	router := gin.Default()
	api.Router(router.Group(env.APIVersion))

	if err := router.Run(fmt.Sprintf(":%s", env.APIPort)); err != nil {
		log.Panicln(err)
	}
}
