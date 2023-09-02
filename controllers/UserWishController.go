package controllers

import (
	"net/http"
	"r101/models"
	"r101/models/dto"
	"r101/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ListWish(c echo.Context) error {
	user_id, err := utils.GetUserID(c)

	if(err != nil){
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err,
		})
	}
	offset := c.QueryParam("offset")
	limit := c.QueryParam("limit")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	offset_int, _ := strconv.ParseInt(offset, 0, 64)
	limit_int, _ := strconv.ParseInt(limit, 0, 64)

	var userWishes []dto.UserWishList

	result := utils.DB.Model(&models.UserWish{}).Where("user_id = ?", user_id).Order("created_at desc").Offset(int(offset_int)).Limit(int(limit_int)).Find(&userWishes)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrive user wishes",
		})
	}

	return c.JSON(http.StatusOK, dto.BaseResponse {
		Status: true,
		Message: "Successfully retrive user wishes.",
		Data: echo.Map{
			"user_wishes" : userWishes,
		},
	})
}

func CreateWish(c echo.Context) error {
	user_id, err := utils.GetUserID(c)

	if(err != nil){
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	body := new(models.UserWish)
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

	userWish := body
	userWish.UserID = user_id
	
	if err := utils.DB.Create(&userWish).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create wishlist",
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.BaseResponse {
		Status: true,
		Message: "Successfully create wish.",
		Data: echo.Map{
			"id" : userWish.ID,
		},
	})
}

func GetWish(c echo.Context) error {
	user_id, err := utils.GetUserID(c)

	if(err != nil){
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Unauthorized!",
		})
	}

	id := c.Param("id")

	var userWish dto.UserWishList
	if err := utils.DB.Model(&models.UserWish{}).Where("user_id = ?", user_id).Where("id = ?", id).First(&userWish).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Data not found",
		})
	}

	return c.JSON(http.StatusOK, dto.BaseResponse {
		Status: true,
		Message: "Successfully retrive user wish.",
		Data:  userWish,
	})
}

func UpdateWish(c echo.Context) error {
	user_id, err := utils.GetUserID(c)

	if(err != nil){
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Unauthorized!",
		})
	}

	updatedData := new(models.UserWish)
	if err := c.Bind(updatedData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid JSON format",
		})
	}

	listError := utils.CustomValidation(updatedData)
	
	if listError != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Errors validation",
			"errors":  listError,
		})
	}

	id := c.Param("id")

	var existingData models.UserWish
	if err := utils.DB.Model(&models.UserWish{}).Where("user_id = ?", user_id).First(&existingData, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Data not found",
		})
	}

	existingData.Name = updatedData.Name
	existingData.Price = updatedData.Price
	existingData.Location = updatedData.Location

	if err := utils.DB.Save(&existingData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update wish",
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.BaseResponse {
		Status: true,
		Message: "Successfully update wish.",
		Data: echo.Map{
			"id" : existingData.ID,
		},
	})
}

func DeleteWish(c echo.Context) error {
	user_id, err := utils.GetUserID(c)

	if(err != nil){
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Unauthorized!",
		})
	}

	id := c.Param("id")

	var existingData models.UserWish
	if err := utils.DB.Model(&models.UserWish{}).Where("user_id = ?", user_id).First(&existingData, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Data not found",
		})
	}

	// Soft delete
	if err := utils.DB.Model(&models.UserWish{}).Delete(&existingData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete wish",
			"errors": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateStatusWish(c echo.Context) error {
	user_id, err := utils.GetUserID(c)

	if(err != nil){
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Unauthorized!",
		})
	}

	updatedData := new(dto.UserWishStatus)
	if err := c.Bind(updatedData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid JSON format",
		})
	}

	listError := utils.CustomValidation(updatedData)
	
	if listError != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Errors validation",
			"errors":  listError,
		})
	}

	id := c.Param("id")

	var existingData models.UserWish
	if err := utils.DB.Model(&models.UserWish{}).Where("user_id = ?", user_id).First(&existingData, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Data not found",
		})
	}

	existingData.IsDone = updatedData.IsDone

	if err := utils.DB.Save(&existingData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update wish",
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.BaseResponse {
		Status: true,
		Message: "Successfully update wish.",
		Data: echo.Map{
			"id" : existingData.ID,
		},
	})
}