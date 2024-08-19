package handler

import (
	"fmt"

	"github.com/dwprz/prasorganic-notification-service/src/interface/service"
	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
	"github.com/gofiber/fiber/v2"
)

type Notification struct {
	notifService service.Notification
}

func NewNotification(ns service.Notification) *Notification {
	return &Notification{
		notifService: ns,
	}
}

func (n *Notification) Midtrans(c *fiber.Ctx) error {
	tx, ok := c.Locals("midtrans").(*entity.Transaction)
	if !ok {
		return fmt.Errorf("unexpected type %T", tx)
	}

	n.notifService.Midtrans(c.Context(), tx)

	return c.Status(200).JSON(fiber.Map{"data": "success"})
}

func (n *Notification) Shipper(c *fiber.Ctx) error {
	req := new(entity.Shipper)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	n.notifService.Shipper(c.Context(), req)

	return c.Status(200).JSON(fiber.Map{"data": "success"})
}
