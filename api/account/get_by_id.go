package account

import (
	"github.com/gin-gonic/gin"
	model "github.com/lffranca/api/repository/account"
)

func GetByID(c *gin.Context) {
	id := c.Param("id")

	acc, errAcc := model.Get()
	if errAcc != nil {
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	accountsItem, errAccountsItem := acc.GetByID(c.Request.Context(), id)
	if errAccountsItem != nil {
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	c.JSON(200, accountsItem)
}
