package handlers

import (
	"log"

	"github.com/rof20004/healthy-api/application/errors"

	"github.com/gofiber/fiber/v2"
)

type errorResponse struct {
	Message string `json:"errorMessage"`
	Code    int    `json:"code"`
}

type successResponse struct {
	Data interface{} `json:"data"`
}

func SendResponse(c *fiber.Ctx, data interface{}, code int, err error) error {
	switch t := err.(type) {
	case errors.CustomError:
		if t.Cause != nil {
			log.Println(t.Cause)
		}

		return c.Status(t.Code).JSON(errorResponse{t.Message, t.Code})
	case error:
		if t != nil {
			log.Println(t.Error())
		}

		var (
			message = "Unexpected error"
			status  = fiber.StatusInternalServerError
		)

		return c.Status(status).JSON(errorResponse{message, status})
	default:
		return c.Status(code).JSON(successResponse{data})
	}
}
