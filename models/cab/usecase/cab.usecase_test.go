package usecase

import (
	geo "github.com/kellydunn/golang-geo"
	"github.com/stretchr/testify/assert"
	"testing"
	"wednesday/container"
	"wednesday/dtos"
)

var ride  = 0
var drive = 0

func Test_useCaseCab_BookCab(t *testing.T) {
	di := container.Container{Profile: "local"}
	di.TriggerDI()
	src := geo.NewPoint(18.5616536,73.8817828)
	dst := geo.NewPoint(18.4676107,73.8997356)
	Test_useCaseCab_CreateRider(t)
	Test_useCaseCab_CreateDriver(t)
	input := dtos.BookCabDto{
		Rider:       ride,
		Driver:      drive,
		Source:      *src,
		Destination: *dst,
		PaymentMode: "cod",
	}
	err := di.GetCabUseCase().BookCab(&input,nil)
	assert.Equal(t,nil,err)
}

func Test_useCaseCab_CancelRide(t *testing.T) {

}

func Test_useCaseCab_CompleteRide(t *testing.T) {

}

func Test_useCaseCab_CreateDriver(t *testing.T) {
   di := container.Container{Profile: "local"}
   di.TriggerDI()
   input := dtos.DriverDto{
	   Name:    "Prateek Tyagi",
	   Mobile:  "7060404050",
	   Vehicle: "DLSM113377",
   }
	id,err := di.GetCabUseCase().CreateDriver(&input,nil)
	if id != 0 {
		drive=id
	}
   assert.Equal(t,nil,err)
	assert.NotEqual(t,0,id)
}

func Test_useCaseCab_CreateRider(t *testing.T) {
   di := container.Container{Profile: "local"}
   di.TriggerDI()
   input := dtos.RiderDto{
	   Name:   "Aavyay Tyagi",
	   Mobile: "7060404050",
	   Email:  "smstoprateek@gmail.com",
   }
   id,err := di.GetCabUseCase().CreateRider(&input,nil)
   if id != 0 {
   	 ride=id
   }
   assert.Equal(t,nil,err)
	assert.NotEqual(t,0,id)
}

func Test_useCaseCab_FeedBack(t *testing.T) {

}

func Test_useCaseCab_GetCompletedRides(t *testing.T) {

}

func Test_useCaseCab_NearByCabs(t *testing.T) {

}

func Test_useCaseCab_StartRide(t *testing.T) {

}
