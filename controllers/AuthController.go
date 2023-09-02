package controllers

import (
	"net/http"

	"r101/models"
	"r101/models/dto"
	"r101/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
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

	return c.JSON(http.StatusCreated, dto.BaseResponse {
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

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
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