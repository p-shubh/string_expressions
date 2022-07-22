package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

	str1 := strings.Split(reqBody.Data, "")

	D, _ := fmt.Println("str1", str1)

	a := str1[0]
	b := str1[1]
	d := str1[2]

	fmt.Println("abc", a, b, d)

	res := gin.H{
		// "error":  err.Error(),
		"result": D,
	}
	c.JSON(http.StatusOK, res)

	strInt, _ := strconv.Atoi(str1[0])

	// for i := 0; i < len(str1); i++ {
	// 	if str1[i]=="+" {
	// 		"+"= +
	// 	}
	// }

	var retur bool = true

	for i := 0; i <= 9; i++ {
		if strInt == i {
			retur = true
		}
	}

	fmt.Println("retur :", retur, str1)

}
