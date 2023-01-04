package db

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init(dialector gorm.Dialector, opts gorm.Option) {
	var err error
	if DB == nil {
		DB, err = gorm.Open(dialector, opts)
		if err != nil {
			panic(err)
		}
	}
}
