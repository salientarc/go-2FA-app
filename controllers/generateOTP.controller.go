package controllers

import (
	"go-2fa-app/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"

	"github.com/joho/godotenv"
)

func (ac *AuthController) GenerateOTP(ctx *gin.Context) {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	var payload *models.OTPInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		return
	}

	issuerDomain := os.Getenv("ISSUERDOMAIN")
	issuerAccountName := os.Getenv("ACCOUNTNAME")

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuerDomain,
		AccountName: issuerAccountName,
		SecretSize:  15,
	})

	if err != nil {
		panic(err)
	}

	var user models.User
	result := ac.DB.First(&user, "id = ?", payload.UserId)
	if result.Error != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Invalid email or Password",
			})
		return
	}

	dataToUpdate := models.User{
		Otp_secret:   key.Secret(),
		Otp_auth_url: key.URL(),
	}

	ac.DB.Model(&user).Updates(dataToUpdate)

	otpResponse := gin.H{
		"base32":      key.Secret(),
		"otpauth_url": key.URL(),
	}

	ctx.JSON(http.StatusOK, otpResponse)
}
