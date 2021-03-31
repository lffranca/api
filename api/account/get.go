package account

import (
	"github.com/gin-gonic/gin"
	model "github.com/lffranca/api/repository/account"
)

func Get(ctx *gin.Context) {
	acc, errAcc := model.Get()
	if errAcc != nil {
		ctx.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	accounts, errAccounts := acc.Get(ctx.Request.Context())
	if errAccounts != nil {
		ctx.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	ctx.JSON(200, accounts)
}
