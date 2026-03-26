package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BodyInvalidMsg  = "body invalid"
	DataNotFoundMsg = "data not found"
)

func Success(c *gin.Context, data ...any) {
	res := gin.H{"success": true}

	if len(data) > 0 {
		res["data"] = data
	}

	c.JSON(http.StatusOK, res)
}

func Error(c *gin.Context, status int, msg *string, data ...any) {
	res := gin.H{"success": false}

	if msg != nil {
		res["msg"] = msg
	}

	if len(data) > 0 {
		res["data"] = data
	}

	c.JSON(status, data)
}
