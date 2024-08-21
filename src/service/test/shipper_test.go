package test

import (
	"context"
	"testing"
	"time"

	"github.com/dwprz/prasorganic-notification-service/src/core/broker/producer"
	"github.com/dwprz/prasorganic-notification-service/src/interface/service"
	"github.com/dwprz/prasorganic-notification-service/src/mock/delivery"
	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
	serviceimpl "github.com/dwprz/prasorganic-notification-service/src/service"
	"github.com/stretchr/testify/suite"
)

type ShipperTestSuite struct {
	suite.Suite
	notifService service.Notification
}

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_Shipper$ -v ./src/service/test/ -count=1

func (s *ShipperTestSuite) SetupSuite() {
	midtransDelivery := delivery.NewMidtransBrokerMock()
	shipperDelivery := delivery.NewShipperBrokerMock()

	kafkaProducer := producer.NewKafka(midtransDelivery, shipperDelivery)

	s.notifService = serviceimpl.NewNotification(kafkaProducer)
}

func (s *ShipperTestSuite) Test_Success() {
	req := s.CreateShipperReq()

	s.notifService.Shipper(context.Background(), req)
}

func (s *ShipperTestSuite) Test_WithoutShippingId() {
	req := s.CreateShipperReq()
	req.ShippingId = ""

	s.notifService.Shipper(context.Background(), req)
}

func (s *ShipperTestSuite) CreateShipperReq() *entity.Shipper {
	return &entity.Shipper{
		Auth:            "Bearer abcdef123456",
		ShippingId:      "ship-001",
		TrackingId:      "tracking-123456",
		OrderTrackingId: "order-track-7890",
		OrderId:         "order-001",
		StatusDate:      time.Now(),
		Internal: entity.InternalExternal{
			Id:          1,
			Name:        "External Courier",
			Description: "Courier service for internal orders",
		},
		External: entity.InternalExternal{
			Id:          2,
			Name:        "External Courier",
			Description: "Third-party courier service",
		},
		InternalStatus: entity.Status{
			Code:        300,
			Name:        "In Transit",
			Description: "Order has been processed internally",
		},
		ExternalStatus: entity.Status{
			Code:        200,
			Name:        "Processed",
			Description: "Order is on the way to the destination",
		},
		AWB: "AWB1234567890",
	}
}

func TestService_Shipper(t *testing.T) {
	suite.Run(t, new(ShipperTestSuite))
}
