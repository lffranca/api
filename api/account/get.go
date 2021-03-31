package account

import (
	"github.com/gin-gonic/gin"
	model "github.com/lffranca/api/repository/account"
)

func Get(c *gin.Context) {
	acc, errAcc := model.Get()
	if errAcc != nil {
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	accounts, errAccounts := acc.Get(c.Request.Context())
	if errAccounts != nil {
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	c.JSON(200, accounts)
}
