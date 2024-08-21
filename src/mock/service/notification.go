package service

import (
	"context"

	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
)

type Notification interface {
	Midtrans(ctx context.Context, data *entity.Transaction)
	Shipper(ctx context.Context, data *entity.Shipper)
}
