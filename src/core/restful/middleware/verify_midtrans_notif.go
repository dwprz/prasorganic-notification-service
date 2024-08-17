package middleware

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/dwprz/prasorganic-notification-service/src/common/errors"
	"github.com/dwprz/prasorganic-notification-service/src/infrastructure/config"
	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) VerifyMidtransNotif(c *fiber.Ctx) error {

	tx := new(entity.Transaction)
	if err := c.BodyParser(&tx); err != nil {
		return err
	}

	key := tx.OrderId + tx.StatusCode + tx.GrossAmount + config.Conf.Midtrans.ServerKey

	hash := sha512.New()
	hash.Write([]byte(key))

	hashByte := hash.Sum(nil)
	hashHex := hex.EncodeToString(hashByte)

	if tx.SignatureKey != hashHex {
		return &errors.Response{HttpCode: 401, Message: "invalid signature key"}
	}

	c.Locals("midtrans", tx)

	return c.Next()
}
