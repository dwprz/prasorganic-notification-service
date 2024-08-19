package service

import (
	"context"

	"github.com/dwprz/prasorganic-notification-service/src/common/log"
	"github.com/dwprz/prasorganic-notification-service/src/core/broker/producer"
	v "github.com/dwprz/prasorganic-notification-service/src/infrastructure/validator"
	"github.com/dwprz/prasorganic-notification-service/src/interface/service"
	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
	"github.com/sirupsen/logrus"
)

type NotificationImpl struct {
	kafkaClient *producer.Kafka
}

func NewNotification(kc *producer.Kafka) service.Notification {
	return &NotificationImpl{
		kafkaClient: kc,
	}
}

func (n *NotificationImpl) Midtrans(ctx context.Context, data *entity.Transaction) {
	if err := v.Validate.Struct(data); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "service.NotificationImpl/Midtrans", "section": "Validate.Struct"}).Error(err)
		return
	}

	n.kafkaClient.Midtrans.Publish(data)
}

func (n *NotificationImpl) Shipper(ctx context.Context, data *entity.Shipper) {
	if err := v.Validate.Struct(data); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "service.NotificationImpl/Shipper", "section": "Validate.Struct"}).Error(err)
		return
	}

	n.kafkaClient.Shipper.Publish(data)
}
