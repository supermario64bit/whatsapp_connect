package types

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApplicationError struct {
	HttpStatus int
	Message    string
	Err        error
}

func (err *ApplicationError) WriteHttpResponse(c *gin.Context) {
	if err.HttpStatus == 0 || err.Message == "" {
		c.JSON(http.StatusNoContent, gin.H{"message": "No content"})
		return
	}

	var resp gin.H
	if err == nil {
		resp = gin.H{
			"status":  "failed",
			"message": err.Message,
		}
	} else {
		resp = gin.H{
			"status":  "failed",
			"message": err.Message,
			"result":  gin.H{"error": err.Err.Error()},
		}
	}

	c.JSON(err.HttpStatus, resp)
}
