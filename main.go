package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectMongo()
	router := gin.New()
	router.POST("/user/register", RegisterUser)
	router.Run(":8080")
}
