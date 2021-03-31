package transaction

import (
	"github.com/gin-gonic/gin"
	model "github.com/lffranca/api/repository/transaction"
)

func GetByAccountID(c *gin.Context) {
	id := c.Param("id")

	md, errMD := model.Get()
	if errMD != nil {
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	items, errItems := md.GetByAccountID(c.Request.Context(), id)
	if errItems != nil {
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	c.JSON(200, items)
}
