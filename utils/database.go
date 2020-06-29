package utils

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"wednesday/config"
	"wednesday/constants"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

var client *gorm.DB

func StartUp(dsConfig *config.DataSource) (*gorm.DB, error) {

	log.Println("Starting a new DB connection.")
	db, err := CreateDBConnection(dsConfig)
	if err != nil {
		return nil, err
	}

	return db, err
}

//CreateDBConnection godocs
func CreateDBConnection(ds *config.DataSource) (*gorm.DB, error) {
	//dsn := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	ds.Host, ds.Port, ds.Username, ds.Password, ds.DatabaseName)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", ds.Username, ds.Password, ds.Host, ds.Port, ds.DatabaseName)
   log.Printf("DNS: {%+v}",ds)
	var err error
	client, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	client.SingularTable(true)
	log.Println(constants.DBConnectionMSG)
	ApplyDBChangeLog(dsn)
	return client, nil
}

func ApplyDBChangeLog(dsn string) {
	sqlDB, err := sql.Open(constants.DriverName, dsn)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(constants.DBMigrationStart)
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	err = goose.Up(sqlDB, basePath+constants.DbChangeLogDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(constants.DBMigrationEnd)
}
