package handlers

import (
	"task_management_service/database"
	"task_management_service/utils"

	"log"
	"task_management_service/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllTasks(c *fiber.Ctx) error {
	result := []models.Task{}
	db := database.DBConn

	// รับค่าจาก query params
	search := c.Query("search")
	status := c.Query("status")

	query := db.Model(result)

	if search != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Order("created_at desc").Find(&result).Error

	response := utils.ResponseSuccess(c, result, fiber.StatusOK, "OK", "Fetch all data successfully.")
	if err != nil {
		log.Print(err)
		response = utils.ResponseError(c, fiber.StatusNotFound, "NOK", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	return response
}

func GetTaskById(c *fiber.Ctx) error {
	id := c.Params("id")
	result := models.Task{}
	db := database.DBConn
	err := db.Where("id = ?", id).First(&result).Error
	response := utils.ResponseSuccess(c, result, fiber.StatusOK, "OK", "Fetch data by id successfully.")

	if err != nil {
		log.Print(err)
		response = utils.ResponseError(c, fiber.StatusNotFound, "NOK", err.Error())
		return response
	}

	return response
}

func CreateTask(c *fiber.Ctx) error {

	result := models.Task{}

	db := database.DBConn
	err := c.BodyParser(&result)
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusInternalServerError, "NOK", err.Error())
		return response
	}

	validate := validator.New()
	err = validate.Struct(result)
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusUnprocessableEntity, "NOK", err.Error())
		return response
	}

	err = db.Create(&result).Error
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusNotFound, "NOK", err.Error())
		return response

	}

	response := utils.ResponseSuccess(c, result, fiber.StatusCreated, "OK", "Create post successfully.")
	return response

}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	result := models.Task{}
	db := database.DBConn
	err := db.Where("id = ?", id).First(&result).Error

	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusNotFound, "NOK", err.Error())
		return response
	}

	// Parsing body ของ request
	err = c.BodyParser(&result)
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusInternalServerError, "NOK", err.Error())
		return response
	}

	// อัปเดต Task
	err = db.Model(&result).Where("id = ?", id).Updates(&result).Error
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusNotFound, "NOK", err.Error())
		return response
	}

	response := utils.ResponseSuccess(c, result, fiber.StatusOK, "OK", "Update post successfully.")
	return response

}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	result := models.Task{}
	db := database.DBConn
	if err := db.First(&result, "id = ?", id).Error; err != nil {
		log.Print(err)
		return utils.ResponseError(c, fiber.StatusNotFound, "NOK", err.Error())
	}
	if err := db.Delete(&result).Error; err != nil {
		log.Print(err)
		return utils.ResponseError(c, fiber.StatusInternalServerError, "NOK", err.Error())
	}
	return utils.ResponseSuccess(c, fiber.Map{}, fiber.StatusOK, "OK", "Delete task successfully.")
}
