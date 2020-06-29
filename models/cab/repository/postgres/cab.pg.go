package postgres

import (
	"github.com/jinzhu/gorm"
	geo "github.com/paulmach/go.geo"
	"log"
	"strconv"
	"wednesday/constants"
	"wednesday/models"
)

type cabRepository struct {
}

func (c cabRepository) NearByCabs(points *geo.Point,tx *gorm.DB) ([]models.Cab,error) {
	panic("implement me")
}

func (c cabRepository) BookCab(input *models.Cab ,tx *gorm.DB) error {
	log.Printf("Cab Input is: {%+v} \n",input)
	query := getCabQuery(input)
	tx.LogMode(true)
	db := tx.Debug().Exec(query)
	db.Model(&models.Driver{}).Where("id = ?",input.Driver).Update("occupied",true)
	riderquery := getDriverRiderUpdateQuery(input.Source,input.Rider)
	tx.Debug().Exec(riderquery)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func getCabQuery(values *models.Cab) string {
	return  constants.CreateCabQuery+strconv.Itoa((*values).Driver)+","+strconv.Itoa((*values).Rider)+","+strconv.FormatFloat((*values).TotalDistance, 'f', -1, 64)+","+strconv.FormatFloat((*values).TotalAmount, 'f', -1, 64)+",'"+(*values).Status+"','2020-06-28 15:47:53','"+(*values).PaymentMode+"',"+"point("+strconv.FormatFloat((*values).Source.Lat(), 'f', -1, 64)+","+strconv.FormatFloat((*values).Source.Lng(), 'f', -1, 64)+")"+","+"point("+strconv.FormatFloat((*values).Destination.Lat(), 'f', -1, 64)+","+strconv.FormatFloat((*values).Destination.Lng(), 'f', -1, 64)+")"+")"
}

func getDriverRiderUpdateQuery(point geo.Point,id int) string {
	return "update rider set current_location=point("+strconv.FormatFloat(point.Lat(), 'f', -1, 64)+","+strconv.FormatFloat(point.Lng(), 'f', -1, 64)+") where id = "+strconv.Itoa(id)
}

func (c cabRepository)  UpdateOTP(id,startOTP, endOTP int,tx *gorm.DB) error {
	db := tx.Model(&models.Rider{}).Where("id = ?",id).
		Update(map[string]interface{}{
			"start_otp": startOTP,
			"end_otp": endOTP})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
func (c cabRepository) GetCompletedRides(id int,tx *gorm.DB) ([]models.Cab,error) {
	var cabs []models.Cab
	query := "select * from cab c inner join driver r on c.rider=r.id where c.status='completed' and r.id="+strconv.Itoa(id)
	tx.LogMode(true)
	db := tx.Debug().Raw(query).Scan(&cabs)
	if db.Error != nil {
		return nil,db.Error
	}
	return cabs,nil
}

func (c cabRepository) StartRide(rider ,driver int,tx *gorm.DB) {
	query := "update cab set status = 'ongoing' where rider = "+strconv.Itoa(rider)+" AND driver = "+strconv.Itoa(driver)
	tx.LogMode(true)
	db := tx.Debug().Exec(query)
	if db.Error != nil {
		return
	}
}

func (c cabRepository) CompleteRide(rider,driver int,tx *gorm.DB) {
	cabQuery := "update cab set status = 'completed' where rider = "+strconv.Itoa(rider)+" AND driver = "+strconv.Itoa(driver)
	tx.LogMode(true)
	db := tx.Debug().Exec(cabQuery)
	riderQuery := "update rider set end_otp=0,start_otp=0 where id = "+strconv.Itoa(rider)
	db.Exec(riderQuery)
	driverQuery := "update driver set occupied=false where id = "+strconv.Itoa(driver)
	db.Exec(driverQuery)
	if db.Error != nil {
		return
	}
}


func (c cabRepository) FeedBack(rider,driver,rating int,tx *gorm.DB) error {
	ratingRes := []int{}

	if driver != 0 {
		query := "select rating from driver where id = "+strconv.Itoa(driver)
		db := tx.Debug().Raw(query).Scan(&ratingRes)
		if  db.Error != nil {
			return db.Error
		}
		average := (ratingRes[0] + rating)/2
		if ratingRes[0] == 0 {
			average = rating
		}
		upQuery := "update driver set rating = "+strconv.Itoa(average)+" where id = "+strconv.Itoa(driver)
		tx.LogMode(true)
		tx.Debug().Exec(upQuery)
		if db.Error != nil {
			return db.Error
		}
	}
	return nil
}

func (c cabRepository) GetOTP(id int,which string,tx *gorm.DB) (int,error){
	response := new(models.Rider)
	db := tx.Model(&models.Rider{}).Where("id = ?",id).Find(&response)
	if db.Error != nil {
		return 0,db.Error
	}
	if which == "start" {
		return response.StartOTP,nil
	}
	return response.EndOTP,nil
}

func (c cabRepository) CreateRider(input *models.Rider,tx *gorm.DB) (int,error) {
	tx.LogMode(true)
	db := tx.Debug().Create(input)
	if db.Error != nil {
		return 0,db.Error
	}
	return input.ID,nil
}

func (c cabRepository) CreateDriver(input *models.Driver,tx *gorm.DB) (int,error) {
	tx.LogMode(true)
	db := tx.Debug().Create(input)
	if db.Error != nil {
		return 0,db.Error
	}
	return input.ID,nil
}


func NewCabRepository() *cabRepository {
	return &cabRepository{}
}