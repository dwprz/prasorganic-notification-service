package broker

import (
	"github.com/dwprz/prasorganic-notification-service/src/core/broker/delivery"
	"github.com/dwprz/prasorganic-notification-service/src/core/broker/producer"
)

func InitKafkaProducer() *producer.Kafka {
	midtransDelivery := delivery.NewMidtransBroker()
	shipperDelivery := delivery.NewShipperBroker()

	kafkaProducer := producer.NewKafka(midtransDelivery, shipperDelivery)
	return kafkaProducer
}
