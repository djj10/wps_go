package controllers

import (
	"net/http"
	"wps_go/user_service/models"
	"wps_go/user_service/database"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "registration failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func Login(c *gin.Context) {
	// 登录逻辑
	c.JSON(http.StatusOK, gin.H{"message": "login success"})
}

func GetUserInfo(c *gin.Context) {
	// 获取用户信息逻辑
	c.JSON(http.StatusOK, gin.H{"message": "user info"})
}