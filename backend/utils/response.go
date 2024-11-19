package utils

import (
	"task_management_service/types"

	"github.com/gofiber/fiber/v2"
)

func ResponseSuccess(c *fiber.Ctx, result interface{}, status int, statusText string, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"result":     result,
		"status":     status,
		"statusText": statusText,
		"message":    message,
		"error":      false,
	})
}

func ResponseError(c *fiber.Ctx, status int, statusText string, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"result":     nil,
		"status":     status,
		"statusText": statusText,
		"message":    message,
		"error":      true,
	})
}

func ResponsePagination(c *fiber.Ctx, result interface{}, status int, statusText string, message string, meta types.MetaResponse, error bool) error {
	return c.Status(status).JSON(fiber.Map{
		"result":     result,
		"status":     status,
		"statusText": statusText,
		"message":    message,
		//    meta: {
		//     page,
		//     page_size: pageSize,
		//     total_records: totalRecords,
		//   },
		"meta":  meta,
		"error": error,
	})
}
