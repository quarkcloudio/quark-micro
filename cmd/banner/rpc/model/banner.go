package model

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

// 模型
type Banner struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CategoryId int            `json:"category_id" gorm:"size:11;default:0;"`
	Title      string         `json:"title" gorm:"size:200;not null"`
	UrlType    int            `json:"url_type" gorm:"size:1;not null;default:1"`
	Url        string         `json:"url" gorm:"size:200;not null"`
	Status     int            `json:"status" gorm:"size:1;not null;default:1"`
	CoverId    string         `json:"cover_id" gorm:"size:1000;default:null"`
	Sort       int            `json:"sort" gorm:"size:11;default:0;"`
	Deadline   time.Time      `json:"deadline"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

// 获取列表
func (model *Banner) List(position string, order string) (items []*Banner, total int64, err error) {
	var (
		getItems []*Banner
		getTotal int64
	)

	// 查询对象
	query := db.Client.
		Model(&model).
		Select("banners.*").
		Joins("left join banner_categories on banner_categories.id = banners.category_id").
		Where("banners.status", 1)

	if position != "" {
		query.Where("banner_categories.name = ?", position)
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
