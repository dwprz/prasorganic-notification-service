package delivery

import "github.com/dwprz/prasorganic-notification-service/src/model/entity"

type MidtransBroker interface {
	Publish(data *entity.Transaction)
	Close()
}
