package usecase

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	geo "github.com/paulmach/go.geo"
	"io"
	"strconv"
	"time"
	"wednesday/dtos"
	"wednesday/models"
	"wednesday/models/pg_utils"
)

type useCaseCab struct {
	cabRepo     models.CabRepo
	transaction pg_utils.Transaction
}

func NewUseCaseCab(cabRepo models.CabRepo, transaction pg_utils.Transaction) *useCaseCab {
	return &useCaseCab{cabRepo: cabRepo, transaction: transaction}
}

func (o useCaseCab) NearByCabs(tx *gorm.DB, source *dtos.LatAndLong) ([]dtos.CabDtoResponse, error) {
	var response []dtos.CabDtoResponse
	tx, isCreated := o.transaction.GetTxn(tx)
	points := new(geo.Point)
	_ = copier.Copy(points, source)
	result, err := o.cabRepo.NearByCabs(points, tx)
	if err != nil {
		if isCreated {
			o.transaction.Rollback(tx)
			return nil, err
		}
	}
	_ = copier.Copy(&response, &result)
	if isCreated {
		o.transaction.Commit(tx)
	}
	return response, nil
}

func (o useCaseCab) BookCab(request *dtos.BookCabDto, tx *gorm.DB) error {
	tx, isCreated := o.transaction.GetTxn(tx)
	err := o.bookCab(request,tx)
	if err != nil {
		if isCreated {
			o.transaction.Rollback(tx)
		}
		return err
	}
	if isCreated {
		o.transaction.Commit(tx)
	}
	return nil
}

func (o useCaseCab) bookCab(request *dtos.BookCabDto, tx *gorm.DB) error {
	totalDistance := request.Source.DistanceFrom(&request.Destination)
	input := models.Cab{
		Driver:        request.Driver,
		Rider:         request.Rider,
		TotalDistance: Round(totalDistance),
		TotalAmount:   Round(totalDistance * 6),
		Status:        "booked",
		Source:        request.Source,
		Destination:   request.Destination,
		DateOfTravel:  time.Now(),
		PaymentMode:   request.PaymentMode,
	}
	err := o.cabRepo.BookCab(&input, tx)
	if err != nil {
		return err
	}
	err = o.saveOTP(request.Rider,tx)
	if err != nil {
		return err
	}
	return nil
}

func (o useCaseCab) saveOTP(id int, tx *gorm.DB) error{
	startOtp := generateOTP(4)
	endOtp := generateOTP(5)
	err := o.cabRepo.UpdateOTP(id,startOtp,endOtp,tx)
	if err != nil {
		return err
	}
	return nil
}

func generateOTP(max int) int {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0','8'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n >= 6 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	otp,_ := strconv.Atoi(string(b))
	return otp
}

func (o useCaseCab) GetCompletedRides(id int, tx *gorm.DB) ([]dtos.CabDtoResponse, error) {
	var response []dtos.CabDtoResponse
	tx, isCreated := o.transaction.GetTxn(tx)
	result, err := o.cabRepo.GetCompletedRides(id, tx)
	if err != nil {
		if isCreated {
			o.transaction.Rollback(tx)
			return nil, err
		}
	}
	_ = copier.Copy(&response, &result)
	if isCreated {
		o.transaction.Commit(tx)
	}
	return response, nil
}

func (o useCaseCab) StartRide(rider ,driver,otp int, tx *gorm.DB) error{
	tx, isCreated := o.transaction.GetTxn(tx)
	startOtp, err := o.cabRepo.GetOTP(rider, "start", tx)
	if err != nil {
		return err
	}
	if otp != startOtp {
		return errors.New("")
	}
	o.cabRepo.StartRide(rider,driver, tx)
	if isCreated {
		o.transaction.Commit(tx)
	}
	return nil
}

func (o useCaseCab) CompleteRide(driver,rider ,otp int, tx *gorm.DB) error{
	tx, isCreated := o.transaction.GetTxn(tx)
	endOTP, err := o.cabRepo.GetOTP(driver, "end", tx)
	if err != nil {
		return err
	}
	if otp != endOTP {
		return errors.New("otp not match try again")
	}
	o.cabRepo.CompleteRide(driver,rider, tx)
	if isCreated {
		o.transaction.Commit(tx)
	}
	return nil
}

func (o useCaseCab) CreateRider(input *dtos.RiderDto,tx *gorm.DB) (int,error) {
	tx,isCreated := o.transaction.GetTxn(tx)
	ride := new(models.Rider)
	_ = copier.Copy(ride, input)
	id,err := o.cabRepo.CreateRider(ride,tx)
	if err != nil {
		if isCreated {
			o.transaction.Rollback(tx)
		}
		return 0,err
	}
	if isCreated {
		o.transaction.Commit(tx)
	}
	return id,nil
}
func (o useCaseCab) CreateDriver(input *dtos.DriverDto ,tx *gorm.DB) (int,error) {
	tx,isCreated := o.transaction.GetTxn(tx)
	driver := new(models.Driver)
	_ = copier.Copy(driver, input)
	id,err := o.cabRepo.CreateDriver(driver,tx)
	if err != nil {
		if isCreated {
			o.transaction.Rollback(tx)
		}
		return 0,err
	}
	if isCreated {
		o.transaction.Commit(tx)
	}
	return id,nil
}

func (o useCaseCab) FeedBack(rider,driver,rating int, tx *gorm.DB) error{
	tx, isCreated := o.transaction.GetTxn(tx)
	err := o.cabRepo.FeedBack(rider,driver, rating, tx)
	if err != nil {
		if isCreated {
			o.transaction.Rollback(tx)
			return err
		}
	}
	if isCreated {
		o.transaction.Commit(tx)
	}
	return nil
}

func Round(rounded float64) float64 {
	formatted, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rounded), 64)
	return formatted
}