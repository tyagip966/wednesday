package pg_utils

import (
	"github.com/jinzhu/gorm"
)

type transaction struct {
	DB *gorm.DB
}

func NewTransaction(db *gorm.DB) Transaction {
	return &transaction{DB: db}
}

func (t *transaction) CreateTransaction() *gorm.DB {
	return t.DB.Begin()
}

func (t *transaction) Commit(tx *gorm.DB) {
	tx.Commit()
}

func (t *transaction) Rollback(tx *gorm.DB) {
	tx.Rollback()
}

func (t *transaction) GetTxn(tx *gorm.DB) (*gorm.DB , bool){
	if tx == nil {
		return t.CreateTransaction(), true
	}
	return tx, false
}