package restful

import (
	"github.com/dwprz/prasorganic-notification-service/src/core/restful/handler"
	"github.com/dwprz/prasorganic-notification-service/src/core/restful/middleware"
	"github.com/dwprz/prasorganic-notification-service/src/core/restful/server"
	"github.com/dwprz/prasorganic-notification-service/src/interface/service"
)

func InitServer(ns service.Notification) *server.Restful {
	notifHandler := handler.NewNotificationRESTful(ns)
	middleware := middleware.New()

	restfulServer := server.NewRestful(notifHandler, middleware)
	return restfulServer
}
