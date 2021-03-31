package transaction

import (
	"github.com/gin-gonic/gin"
	model "github.com/lffranca/api/repository/transaction"
)

func GetByAccountID(ctx *gin.Context) {
	id := ctx.Param("id")

	md, errMD := model.Get()
	if errMD != nil {
		ctx.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	items, errItems := md.GetByAccountID(ctx.Request.Context(), id)
	if errItems != nil {
		ctx.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	ctx.JSON(200, items)
}
