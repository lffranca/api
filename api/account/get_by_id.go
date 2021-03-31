package account

import (
	"github.com/gin-gonic/gin"
	model "github.com/lffranca/api/repository/account"
)

func GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	acc, errAcc := model.Get()
	if errAcc != nil {
		ctx.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	accountsItem, errAccountsItem := acc.GetByID(ctx.Request.Context(), id)
	if errAccountsItem != nil {
		ctx.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	ctx.JSON(200, accountsItem)
}
