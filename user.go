package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	fmt.Println("Received here")
	var user User
	c.BindJSON(&user)
	err := SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
