package models

import (
	"github.com/jinzhu/gorm"
	geo "github.com/paulmach/go.geo"
	"time"
	"wednesday/dtos"
)

type Cab struct {
	ID            int       `json:"id"`
	Driver        int       `json:"driver"`
	Rider         int       `json:"rider"`
	TotalDistance float64   `json:"total_distance"`
	TotalAmount   float64   `json:"total_amount"`
	Status        string    `json:"status"`
	Source        geo.Point `gorm:"Point" json:"source"`
	Destination   geo.Point `gorm:"Point" json:"destination"`
	DateOfTravel  time.Time `json:"date_of_travel"`
	PaymentMode   string    `json:"payment_mode"`
}

type Rider struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Mobile          string    `json:"mobile"`
	Email           string    `json:"email"`
	Rating          int       `json:"rating"`
	StartOTP        int       `json:"start_otp"`
	EndOTP          int       `json:"end_otp"`
	CurrentLocation geo.Point `json:"current_location"`
}

type Driver struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Mobile          string    `json:"mobile"`
	VehicleNo       string    `json:"vehicle_no"`
	Category        string    `json:"category"`
	Occupied        bool      `json:"occupied"`
	Rating          int       `json:"rating"`
	CurrentLocation geo.Point `json:"current_location"`
}

type CabUseCase interface {
	NearByCabs(tx *gorm.DB, source *dtos.LatAndLong) ([]dtos.CabDtoResponse, error)
	BookCab(input *dtos.BookCabDto ,tx *gorm.DB) error
	GetCompletedRides(id int, tx *gorm.DB) ([]dtos.CabDtoResponse, error)
	StartRide(rider ,driver,otp int, tx *gorm.DB) error
	CompleteRide(driver,rider ,otp int, tx *gorm.DB) error
	FeedBack(rider,driver,rating int, tx *gorm.DB) error
	CreateRider(input *dtos.RiderDto,tx *gorm.DB) (int,error)
	CreateDriver(input *dtos.DriverDto ,tx *gorm.DB) (int,error)
}

type CabRepo interface {
	NearByCabs(points *geo.Point, tx *gorm.DB) ([]Cab, error)
	BookCab(input *Cab ,tx *gorm.DB) error
	GetCompletedRides(id int, tx *gorm.DB) ([]Cab, error)
	StartRide(rider ,driver int, tx *gorm.DB)
	CompleteRide(rider,driver int, tx *gorm.DB)
	FeedBack(rider,driver,rating int, tx *gorm.DB) error
    GetOTP(id int,which string,tx *gorm.DB) (int,error)
	UpdateOTP(id,startOTP, endOTP int,tx *gorm.DB) error
	CreateRider(input *Rider,tx *gorm.DB) (int,error)
	CreateDriver(input *Driver,tx *gorm.DB) (int,error)
}
