package dal

import (
	"github.com/quarkcms/quark-hertz/pkg/resource/dal/db"
	"gorm.io/gorm"
)

// Init init DB
func InitDB(dialector gorm.Dialector, opts gorm.Option) {
	db.Init(dialector, opts)
}
