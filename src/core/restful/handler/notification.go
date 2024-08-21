package handler

import (
	"fmt"

	"github.com/dwprz/prasorganic-notification-service/src/interface/service"
	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
	"github.com/gofiber/fiber/v2"
)

type NotificationRESTful struct {
	notifService service.Notification
}

func NewNotificationRESTful(ns service.Notification) *NotificationRESTful {
	return &NotificationRESTful{
		notifService: ns,
	}
}

func (n *NotificationRESTful) Midtrans(c *fiber.Ctx) error {
	tx, ok := c.Locals("midtrans").(*entity.Transaction)
	if !ok {
		return fmt.Errorf("unexpected type %T", tx)
	}

	n.notifService.Midtrans(c.Context(), tx)

	return c.Status(200).JSON(fiber.Map{"data": "success"})
}

func (n *NotificationRESTful) Shipper(c *fiber.Ctx) error {
	req := new(entity.Shipper)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	n.notifService.Shipper(c.Context(), req)

	return c.Status(200).JSON(fiber.Map{"data": "success"})
}
