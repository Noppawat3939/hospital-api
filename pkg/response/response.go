package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BodyInvalidMsg  = "body invalid"
	DataNotFoundMsg = "data not found"
	UnAuthorized    = "user unauthorized"
)

func Success(c *gin.Context, data ...any) {
	res := gin.H{"success": true}

	if len(data) > 0 {
		res["data"] = data[0]
	}

	c.JSON(http.StatusOK, res)
}

func Error(c *gin.Context, status int, msg string, data ...any) {
	res := gin.H{"success": false}

	if msg != "" {
		res["message"] = msg
	}

	if len(data) > 0 {
		res["data"] = data
	}

	c.JSON(status, res)
}
