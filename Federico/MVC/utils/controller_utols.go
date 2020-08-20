package utils

import (
	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	} else {
		c.JSON(status, body)

	}
}

func RespondError(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
		return
	} else {
		c.JSON(err.StatusCode, err)

	}
}
