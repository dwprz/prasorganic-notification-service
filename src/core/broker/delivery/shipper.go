package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/dwprz/prasorganic-notification-service/src/common/log"
	"github.com/dwprz/prasorganic-notification-service/src/infrastructure/config"
	"github.com/dwprz/prasorganic-notification-service/src/interface/delivery"
	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type ShipperBrokerImpl struct {
	writer *kafka.Writer
}

func NewShipperBroker() delivery.Broker[*entity.Shipper] {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(config.Conf.Kafka.Addr1, config.Conf.Kafka.Addr2, config.Conf.Kafka.Addr3),
		Topic:                  "shipper",
		Balancer:               &kafka.LeastBytes{},
		RequiredAcks:           kafka.RequireAll,
		AllowAutoTopicCreation: true,
		WriteTimeout:           time.Duration(10 * time.Second),
		ReadTimeout:            time.Duration(10 * time.Second),

		Logger: kafka.LoggerFunc(func(s string, i ...interface{}) {
			log.Logger.Infof(s, i...)
		}),
	}

	return &ShipperBrokerImpl{
		writer: w,
	}
}

func (s *ShipperBrokerImpl) Publish(data *entity.Shipper) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "delivery.ShipperBrokerImpl/Publish", "section": "json.Marshal"}).Error(err)
		return
	}

	msg := kafka.Message{
		Value: jsonData,
	}

	const maxRetries = 3

	for i := 0; i < maxRetries; i++ {
		err := s.writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Logger.WithFields(logrus.Fields{"location": "delivery.ShipperBrokerImpl/Publish", "section": "writer.WriteMessages"}).Error(err)
		}

		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, kafka.KafkaStorageError) {
			time.Sleep(time.Millisecond * 250) // waktu tunggu antara retries
			continue
		}

		break
	}
}

func (s *ShipperBrokerImpl) Close() {
	if err := s.writer.Close(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "delivery.ShipperBrokerImpl/Close", "section": "writer.Close"}).Error(err)
	}
}
