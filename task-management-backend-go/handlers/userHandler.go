package handlers

import (
	"task_management_service/database"
	"task_management_service/utils"

	"log"
	"task_management_service/models"

	"github.com/gofiber/fiber/v2"
)

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	result := models.User{}
	db := database.DBConn
	err := db.Where("id = ?", id).First(&result).Error

	result.Password = ""
	result.TwoFactorToken = ""
	response := utils.ResponseSuccess(c, result, fiber.StatusOK, "OK", "Fetch data by id successfully.")

	if err != nil {
		log.Print(err)
		response = utils.ResponseError(c, fiber.StatusNotFound, "NOK", err.Error())
		return response
	}

	return response
}
