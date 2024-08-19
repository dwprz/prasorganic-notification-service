package producer

import "github.com/dwprz/prasorganic-notification-service/src/interface/delivery"

type Kafka struct {
	Midtrans delivery.MidtransBroker
	Shipper  delivery.ShipperBroker
}

func NewKafka(md delivery.MidtransBroker, sd delivery.ShipperBroker) *Kafka {
	return &Kafka{
		Midtrans: md,
		Shipper:  sd,
	}
}

func (k *Kafka) Close() {
	k.Midtrans.Close()
	k.Shipper.Close()
}
