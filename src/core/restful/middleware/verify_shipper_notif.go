package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/dwprz/prasorganic-notification-service/src/common/errors"
	"github.com/dwprz/prasorganic-notification-service/src/infrastructure/config"
	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) VerifyShipperNotif(c *fiber.Ctx) error {
	responseFormat := "json"

	data := fmt.Sprintf("%s%s/api/notifications/shipper%s", config.Conf.Shipper.ApiKey, config.Conf.Ngrok.BaseUrl, responseFormat)

	hash := md5.Sum([]byte(data))
	hashString := hex.EncodeToString(hash[:])

	req := new(entity.Shipper)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if hashString != req.Auth {
		return &errors.Response{HttpCode: 401, Message: "invalid authentication"}
	}

	c.Locals("shipper", req)

	return c.Next()
}
