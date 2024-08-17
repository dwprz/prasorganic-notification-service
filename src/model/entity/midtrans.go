package entity

type Transaction struct {
	TransactionTime        string `json:"transaction_time"`
	TransactionStatus      string `json:"transaction_status"`
	TransactionId          string `json:"transaction_id" validate:"required"`
	StatusMessage          string `json:"status_message" validate:"required"`
	StatusCode             string `json:"status_code" validate:"required"`
	SignatureKey           string `json:"signature_key" validate:"required"`
	PaymentType            string `json:"payment_type" validate:"required"`
	OrderId                string `json:"order_id" validate:"required"`
	MerchantId             string `json:"merchant_id"`
	MaskedCard             string `json:"masked_card"`
	GrossAmount            string `json:"gross_amount" validate:"required"`
	FraudStatus            string `json:"fraud_status" validate:"required"`
	Eci                    string `json:"eci"`
	Currency               string `json:"currency"`
	ChannelResponseMessage string `json:"channel_response_message"`
	ChannelResponseCode    string `json:"channel_response_code"`
	CardType               string `json:"card_type"`
	Bank                   string `json:"bank"`
	ApprovalCode           string `json:"approval_code"`
}
