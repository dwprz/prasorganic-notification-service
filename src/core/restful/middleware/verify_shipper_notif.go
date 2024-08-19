package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/dwprz/prasorganic-notification-service/src/common/errors"
	"github.com/dwprz/prasorganic-notification-service/src/infrastructure/config"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) VerifyShipperNotif(c *fiber.Ctx) error {
	responseFormat := "json"

	data := fmt.Sprintf("%s%s/api/notifications/shipper%s", config.Conf.Shipper.ApiKey, config.Conf.Shipper.BaseUrl, responseFormat)

	hash := md5.Sum([]byte(data))
	hashString := hex.EncodeToString(hash[:])

	shippingAuth := c.FormValue("auth")

	if hashString != shippingAuth {
		return &errors.Response{HttpCode: 401, Message: "invalid authentication"}
	}

	return c.Next()
}
