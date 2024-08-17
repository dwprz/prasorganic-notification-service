package producer

import "github.com/dwprz/prasorganic-notification-service/src/interface/delivery"

type Kafka struct {
	Midtrans delivery.MidtransBroker
}

func NewKafka(mb delivery.MidtransBroker) *Kafka {
	return &Kafka{
		Midtrans: mb,
	}
}

func (m *Kafka) Close() {
	m.Midtrans.Close()
}
