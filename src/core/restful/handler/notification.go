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
		return fmt.Errorf("unexpected type %T (*entity.Transaction)", tx)
	}

	n.notifService.Midtrans(c.Context(), tx)

	return c.Status(200).JSON(fiber.Map{"data": "success"})
}

func (n *NotificationRESTful) Shipper(c *fiber.Ctx) error {
	shipper, ok := c.Locals("shipper").(*entity.Shipper)
	if !ok {
		return fmt.Errorf("unexpected type %T (*entity.Shipper)", shipper)
	}

	n.notifService.Shipper(c.Context(), shipper)

	return c.Status(200).JSON(fiber.Map{"data": "success"})
}
