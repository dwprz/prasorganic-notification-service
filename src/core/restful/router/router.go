package router

import (
	"github.com/dwprz/prasorganic-notification-service/src/core/restful/handler"
	"github.com/dwprz/prasorganic-notification-service/src/core/restful/middleware"
	"github.com/gofiber/fiber/v2"
)

func Notification(app *fiber.App, h *handler.NotificationRESTful, m *middleware.Middleware) {
	// all
	app.Add("POST", "/api/notifications/midtrans", m.VerifyMidtransNotif, h.Midtrans)
	app.Add("POST", "/api/notifications/shipper", m.VerifyShipperNotif, h.Shipper)
}