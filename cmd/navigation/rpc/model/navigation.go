package model

import (
	"time"

	"github.com/quarkcloudio/quark-go/v2/pkg/dal/db"
	"gorm.io/gorm"
)

// 模型
type Navigation struct {
	Id        int            `json:"id" gorm:"autoIncrement"`
	Pid       int            `json:"pid"`
	Title     string         `json:"title" gorm:"size:200;not null"`
	CoverId   string         `json:"cover_id" gorm:"size:500;default:null"`
	Sort      int            `json:"sort" gorm:"size:11;default:0;"`
	UrlType   int            `json:"url_type" gorm:"size:1;not null;default:1"`
	Url       string         `json:"url" gorm:"size:200;not null"`
	Status    int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// 获取列表
func (model *Navigation) List(pid int64, order string) (items []*Navigation, total int64, err error) {
	var (
		getItems []*Navigation
		getTotal int64
	)

	// 查询对象
	query := db.Client.
		Model(&model).
		Where("status", 1)

	if pid != 0 {
		query.Where("pid = ?", pid)
	}

	if order != "" {
		query.Order(order)
	}

	// 查询列表
	query.Find(&getItems)

	// 查询总数量
	query.Count(&getTotal)

	// 获取错误信息
	err = query.Error
	if err != nil {
		return
	}

	return getItems, getTotal, err
}
