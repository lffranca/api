package transaction

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/api/export"
	model "github.com/lffranca/api/repository/transaction"
)

func GetByAccountIDPDF(c *gin.Context) {
	id := c.Param("id")

	md, errMD := model.Get()
	if errMD != nil {
		log.Println("errMD: ", errMD)
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	items, errItems := md.GetByAccountID(c.Request.Context(), id)
	if errItems != nil {
		log.Println("errItems: ", errItems)
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	htmlItem, errHTML := export.NewHTML(&export.NewHTMLInput{
		Template: []byte(htmlToPDF),
		Input:    map[string]interface{}{"Transactions": items},
	})
	if errHTML != nil {
		log.Println("errHTML: ", errHTML)
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	pdfItem, errPDFItem := export.NewPDF(&export.NewPDFInput{
		Template: htmlItem.Bytes(),
	})
	if errPDFItem != nil {
		log.Println("errPDFItem: ", errPDFItem)
		c.JSON(400, map[string]string{"message": "Invalid request"})
		return
	}

	readerItem := pdfItem.Reader()

	var contentLength int64 = int64(len(pdfItem.Bytes()))
	contentType := "application/pdf"
	extraHeaders := map[string]string{}

	c.DataFromReader(http.StatusOK, contentLength, contentType, readerItem, extraHeaders)
}

const htmlToPDF = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    <table
      style="
        font-family: arial, sans-serif;
        border-collapse: collapse;
        width: 100%;
      "
    >
      <thead>
        <tr>
          <th
            style="
              border-width: 1px;
              border-style: solid;
              border-color: #dddddd;
              text-align: left;
              padding-top: 8px;
              padding-bottom: 8px;
              padding-right: 8px;
              padding-left: 8px;
            "
          >
            CONTA
          </th>
          <th
            style="
              border-width: 1px;
              border-style: solid;
              border-color: #dddddd;
              text-align: left;
              padding-top: 8px;
              padding-bottom: 8px;
              padding-right: 8px;
              padding-left: 8px;
            "
          >
            TRANSAÇÃO
          </th>
          <th
            style="
              border-width: 1px;
              border-style: solid;
              border-color: #dddddd;
              text-align: left;
              padding-top: 8px;
              padding-bottom: 8px;
              padding-right: 8px;
              padding-left: 8px;
            "
          >
            VALOR
          </th>
        </tr>
      </thead>

      <tbody>
        {{range .Transactions}}
        <tr>
          <td
            style="
              border-width: 1px;
              border-style: solid;
              border-color: #dddddd;
              text-align: left;
              padding-top: 8px;
              padding-bottom: 8px;
              padding-right: 8px;
              padding-left: 8px;
            "
          >
            {{.AccountID}}
          </td>
          <td
            style="
              border-width: 1px;
              border-style: solid;
              border-color: #dddddd;
              text-align: left;
              padding-top: 8px;
              padding-bottom: 8px;
              padding-right: 8px;
              padding-left: 8px;
            "
          >
            {{.Description}}
          </td>
          <td
            style="
              border-width: 1px;
              border-style: solid;
              border-color: #dddddd;
              text-align: left;
              padding-top: 8px;
              padding-bottom: 8px;
              padding-right: 8px;
              padding-left: 8px;
            "
          >
            {{.Value}}
          </td>
        </tr>
        {{end}}
      </tbody>
    </table>
  </body>
</html>

`
