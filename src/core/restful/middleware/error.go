package middleware

import (
	"encoding/json"
	"strconv"

	"github.com/dwprz/prasorganic-notification-service/src/common/errors"
	"github.com/dwprz/prasorganic-notification-service/src/common/errors/restful"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Error(c *fiber.Ctx, err error) error {
	restful.LogError(c, err)

	if validationError, ok := err.(validator.ValidationErrors); ok {
		return restful.HandleValidationError(c, validationError)
	}

	if responseError, ok := err.(*errors.Response); ok {
		return restful.HandleResponseError(c, responseError)
	}

	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		return restful.HandleJsonError(c, jsonError)
	}

	if strconvError, ok := err.(*strconv.NumError); ok {
		return restful.HandleStrconvError(c, strconvError)
	}

	return c.Status(500).JSON(fiber.Map{
		"errors": "sorry, internal server error try again later",
	})
}
