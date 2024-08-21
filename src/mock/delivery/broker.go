package delivery

import (
	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
	"github.com/stretchr/testify/mock"
)

type BrokerMock[T any] struct {
	mock.Mock
}

func NewMidtransBrokerMock() *BrokerMock[*entity.Transaction] {
	return &BrokerMock[*entity.Transaction]{
		Mock: mock.Mock{},
	}
}

func NewShipperBrokerMock() *BrokerMock[*entity.Shipper] {
	return &BrokerMock[*entity.Shipper]{
		Mock: mock.Mock{},
	}
}

func (b *BrokerMock[T]) Publish(data T) {}

func (b *BrokerMock[T]) Close() {}
