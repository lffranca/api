package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lffranca/api/environment"
)

func TestMain(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARNING] Not use .env file")
	}

	env, errENV := environment.Get()
	if errENV != nil {
		log.Panicln(errENV)
	}

	router := gin.Default()
	v1Router := router.Group(env.APIVersion)

	v1Router.GET("/contas", func(c *gin.Context) {})
	v1Router.GET("/contas/:id", func(c *gin.Context) {})
	v1Router.GET("/contas/:id/transacoes", func(c *gin.Context) {})

	if err := router.Run(fmt.Sprintf(":%s", env.APIPort)); err != nil {
		log.Panicln(err)
	}
}
