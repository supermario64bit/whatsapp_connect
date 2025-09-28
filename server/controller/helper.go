package controller

import (
	"github.com/gin-gonic/gin"
)

func writeSuccessHttpResponseObj(message string, resultObjName string, result interface{}) gin.H {
	if result == nil {
		return gin.H{
			"status":  "success",
			"message": message,
		}
	}
	return gin.H{
		"status":  "success",
		"message": message,
		"result": gin.H{
			resultObjName: result,
		},
	}
}

func writeFailedHttpResponseObj(message string, err error) gin.H {
	return gin.H{
		"status":  "failed",
		"message": message,
		"result": gin.H{
			"error": err.Error(),
		},
	}
}
