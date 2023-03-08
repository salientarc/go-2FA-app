package controllers

import (
	"go-2fa-app/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *models.RegisterUserInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		return
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: payload.Password,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(
			http.StatusConflict, gin.H{
				"status":  "fail",
				"message": "Email already exist, please use another email address",
			})
		return
	} else if result.Error != nil {
		ctx.JSON(
			http.StatusBadGateway, gin.H{
				"status":  "error",
				"message": result.Error.Error(),
			})
		return
	}

	ctx.JSON(
		http.StatusCreated, gin.H{
			"status":  "success",
			"message": "Registered successfully, please login",
		})
}
