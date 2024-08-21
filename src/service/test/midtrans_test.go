package test

import (
	"context"
	"testing"

	"github.com/dwprz/prasorganic-notification-service/src/core/broker/producer"
	"github.com/dwprz/prasorganic-notification-service/src/interface/service"
	"github.com/dwprz/prasorganic-notification-service/src/mock/delivery"
	"github.com/dwprz/prasorganic-notification-service/src/model/entity"
	serviceimpl "github.com/dwprz/prasorganic-notification-service/src/service"
	"github.com/stretchr/testify/suite"
)

type MidtransTestSuite struct {
	suite.Suite
	midtransDelivery *delivery.BrokerMock[*entity.Transaction]
	notifService     service.Notification
}

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_Midtrans$ -v ./src/service/test/ -count=1

func (m *MidtransTestSuite) SetupSuite() {
	m.midtransDelivery = delivery.NewMidtransBrokerMock()
	shipperDelivery := delivery.NewShipperBrokerMock()

	kafkaProducer := producer.NewKafka(m.midtransDelivery, shipperDelivery)

	m.notifService = serviceimpl.NewNotification(kafkaProducer)
}

func (m *MidtransTestSuite) Test_Success() {
	req := m.CreateTransactionReq()

	m.notifService.Midtrans(context.Background(), req)
}

func (m *MidtransTestSuite) Test_WithoutSignatureKey() {
	req := m.CreateTransactionReq()
	req.SignatureKey = ""

	m.notifService.Midtrans(context.Background(), req)
}

func (m *MidtransTestSuite) CreateTransactionReq() *entity.Transaction {
	return &entity.Transaction{
		TransactionTime:        "2024-08-21T14:30:00Z",
		TransactionStatus:      "settlement",
		TransactionId:          "abc123xyz789",
		StatusMessage:          "Transaction successful",
		StatusCode:             "200",
		SignatureKey:           "a1b2c3d4e5f6g7h8i9j0",
		PaymentType:            "credit_cart",
		OrderId:                "order-001",
		MerchantId:             "merchant-001",
		MaskedCard:             "4811-xxxx-xxxx-1234",
		GrossAmount:            "50000",
		FraudStatus:            "accept",
		Eci:                    "05",
		Currency:               "IDR",
		ChannelResponseMessage: "Approved",
		ChannelResponseCode:    "00",
		CardType:               "VISA",
		Bank:                   "BCA",
		ApprovalCode:           "12345",
	}
}

func TestService_Midtrans(t *testing.T) {
	suite.Run(t, new(MidtransTestSuite))
}
