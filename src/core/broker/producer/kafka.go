package producer

import (
	"github.com/dwprz/prasorganic-notification-service/src/interface/delivery"
	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
)

type Kafka struct {
	Midtrans delivery.Broker[*entity.Transaction]
	Shipper  delivery.Broker[*entity.Shipper]
}

func NewKafka(midtrans delivery.Broker[*entity.Transaction], shipper delivery.Broker[*entity.Shipper]) *Kafka {
	return &Kafka{
		Midtrans: midtrans,
		Shipper:  shipper,
	}
}

func (k *Kafka) Close() {
	k.Midtrans.Close()
	k.Shipper.Close()
}
