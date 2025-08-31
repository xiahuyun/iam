package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/marmotedu/api/apiserver/v1"
)

// 用户相关处理函数
func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List all users",
		"users":   []interface{}{},
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID",
		"id":      id,
	})
}

// create user handler
func CreateUser(c *gin.Context) {
	// 1. get the user parameters from request body
	var user v1.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. validate the user parameters
	errs := user.Validate()
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Create user",
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Update user",
		"id":      id,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete user",
		"id":      id,
	})
}
