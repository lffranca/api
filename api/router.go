package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/api/api/account"
	"github.com/lffranca/api/api/transaction"
)

func Router(router *gin.RouterGroup) {
	router.GET("/contas", account.Get)
	router.GET("/contas/:id", account.GetByID)
	router.GET("/contas/:id/transacoes", transaction.GetByAccountID)
}
