package transaction

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/api/export"
	model "github.com/lffranca/api/repository/transaction"
)

func GetByAccountIDXLSX(c *gin.Context) {
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

	data := [][]string{{"CONTA", "TRANSAÇÃO", "VALOR"}}
	for _, item := range items {
		row := []string{item.AccountID, item.Description, fmt.Sprint(item.Value)}
		data = append(data, row)
	}

	fileItem, errCSVItem := export.NewXLSX(&export.NewXLSXInput{
		Data: data,
	})
	if errCSVItem != nil {
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	readerItem := fileItem.Reader()

	var contentLength int64 = int64(len(fileItem.Bytes()))
	contentType := "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	extraHeaders := map[string]string{}

	c.DataFromReader(http.StatusOK, contentLength, contentType, readerItem, extraHeaders)
}
