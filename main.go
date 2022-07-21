package main

import "github.com/gin-gonic/gin"

type OpretStrng struct {
	Data string `json:"data"`
}

func main() {

	router := gin.Default()
	setupRoutes(router)
	router.Run(":8080")
}

func setupRoutes(routes *gin.Engine) {

	routes.POST("/expressions", Expressionss)

}
