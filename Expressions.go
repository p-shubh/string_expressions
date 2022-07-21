package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Expressionss(c *gin.Context) {

	reqBody := OpretStrng{}

	c.Bind(&reqBody)

	// if err != nil {
	// 	res := gin.H{
	// 		"error":           err.Error(),
	// 		"result of error": reqBody,
	// 	}

	// 	c.JSON(http.StatusBadRequest, res)
	// }

	if reqBody.Data == "" {
		res := gin.H{
			// "error":  err.Error(),
			"result of reqbody.data": reqBody.Data,
		}

		c.JSON(http.StatusBadRequest, res)
	} else {
		res := gin.H{
			// "error":  err.Error(),
			"result": reqBody.Data,
		}
		c.JSON(http.StatusOK, res)

	}
}
