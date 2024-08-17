package broker

import (
	"github.com/dwprz/prasorganic-notification-service/src/core/broker/producer"
	"github.com/dwprz/prasorganic-notification-service/src/core/broker/delivery"
)

func InitKafkaProducer() *producer.Kafka {
	midtransDelivery := delivery.NewMidtransBroker()
	kafkaProducer := producer.NewKafka(midtransDelivery)

	return kafkaProducer
}
