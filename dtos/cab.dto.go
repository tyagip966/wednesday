package dtos

import (
	geo "github.com/paulmach/go.geo"
)

type CabDtoResponse struct {
	Source        interface{} `json:"source"`
	Destination   interface{} `json:"destination"`
	Driver        interface{} `json:"driver"`
	Rider         interface{} `json:"rider"`
	TotalDistance int         `json:"total_distance"`
	TotalAmount   float64     `json:"total_amount"`
	StartOTP      int         `json:"start_otp"`
	EndOTP        int         `json:"end_otp"`
	PaymentMode   string      `json:"payment_mode"`
}

type LatAndLong struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

type BookCabDto struct {
	Rider       int       `json:"rider"`
	Driver      int       `json:"driver"`
	Source      geo.Point `json:"source"`
	Destination geo.Point `json:"destination"`
	PaymentMode   string `json:"payment_mode"`
}


type RiderDto struct {
	Name string `json:"name"`
	Mobile string `json:"mobile"`
	Email string `json:"email"`
}

type DriverDto struct {
	Name string `json:"name"`
	Mobile string `json:"mobile"`
	VehicleNo string `json:"vehicle_no"`
	Category  string `json:"category"`
}