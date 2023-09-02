package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"prakerja9/models"
	"prakerja9/models/dto"
	"prakerja9/utils"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	body := new(models.User)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Invalid JSON format",
		})
	}

	listError := utils.CustomValidation(body)
	
	if listError != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Errors validation",
			"errors":  listError,
		})
	}

	hasher := sha256.New()

	// Menambahkan data string ke hasher
	_, err := hasher.Write([]byte(body.Password))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to hash password",
		})
	}

	// Mendapatkan hasil hash sebagai byte
	hashBytes := hasher.Sum(nil)

	// Mengonversi hasil hash menjadi string hexadecimal
	hashedPassword := hex.EncodeToString(hashBytes)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to hash password",
		})
	}

	user := models.User{
		Name: 		body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}
	
	if err := utils.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create user",
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.BaseResponse {
		Status: true,
		Message: "Successfully register.",
		Data: echo.Map{
			"email" : user.Email,
		},
	})
}


func Login(c echo.Context) error {
	body := new(dto.UserLogin)
	if err := c.Bind(body); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "Invalid JSON format",
			})
	}

	listError := utils.CustomValidation(body)
	
	if listError != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Errors validation",
			"errors":  listError,
		})
	}

	var user models.User

	if err := utils.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Email not registered",
		})
	}

	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	hasher := sha256.New()

	// Menambahkan data string ke hasher
	_, err := hasher.Write([]byte(body.Password))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to hash password",
		})
	}

	// Mendapatkan hasil hash sebagai byte
	hashBytes := hasher.Sum(nil)

	// Mengonversi hasil hash menjadi string hexadecimal
	hashedPassword := hex.EncodeToString(hashBytes)

	if hashedPassword != user.Password {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid credentials",
		})
	}

	t, err := utils.GenerateJwt(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed generate token",
		})
	}

	return c.JSON(http.StatusOK, dto.BaseResponse {
		Status: true,
		Message: "Successfully signed.",
		Data: echo.Map{
			"token" : t,
		},
	})
}