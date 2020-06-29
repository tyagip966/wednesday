package pg_utils

import "github.com/jinzhu/gorm"

type Transaction interface {
	CreateTransaction() *gorm.DB
	Commit(tx *gorm.DB)
	Rollback(tx *gorm.DB)
	GetTxn(tx *gorm.DB) (*gorm.DB , bool)
}
