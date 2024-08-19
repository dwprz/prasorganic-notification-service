package entity

import (
	"time"
)

type Shipper struct {
	Auth            string           `json:"auth" validate:"required"`
	ShippingId      string           `json:"order_id" validate:"required"` // order_id ini berisi shipping_id dari aplikasi ini
	TrackingId      string           `json:"tracking_id" validate:"required"`
	OrderTrackingId string           `json:"order_tracking_id" validate:"required"`
	OrderId         string           `json:"external_id" validate:"required"` // external_id ini berisi order_id dari aplikasi ini
	StatusDate      time.Time        `json:"status_date" validate:"required"`
	Internal        InternalExternal `json:"internal" validate:"required"`
	External        InternalExternal `json:"external" validate:"required"`
	InternalStatus  Status           `json:"internal_status" validate:"required"`
	ExternalStatus  Status           `json:"external_status" validate:"required"`
	AWB             string           `json:"awb" validate:"required"`
}

type InternalExternal struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type Status struct {
	Code        int    `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
