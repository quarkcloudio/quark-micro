package model

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

// 模型
type Category struct {
	Id          int            `json:"id" gorm:"autoIncrement"`
	Pid         int            `json:"pid"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Sort        int            `json:"sort" gorm:"size:11;default:0;"`
	CoverId     string         `json:"cover_id" gorm:"size:500;default:null"`
	Name        string         `json:"name" gorm:"size:100;default:null"`
	Description string         `json:"description" gorm:"size:500;default:null"`
	Count       int            `json:"count" gorm:"size:11;default:10;"`
	IndexTpl    string         `json:"index_tpl" gorm:"size:100;"`
	ListTpl     string         `json:"list_tpl" gorm:"size:100;"`
	DetailTpl   string         `json:"detail_tpl" gorm:"size:100;"`
	PageNum     int            `json:"page_num" gorm:"size:11;default:10;"`
	Type        string         `json:"type" gorm:"size:200;not null;default:ARTICLE"`
	Status      int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

// 获取详细信息
func (model *Category) Info(id int64, name string) (category *Category, err error) {

	// 查询对象
	query := db.Client.
		Model(&model).
		Where("status", 1)

	if id != 0 {
		query.Where("id = ?", id)
	}

	if name != "" {
		query.Where("name = ?", name)
	}

	// 查询信息
	err = query.Find(&category).Error
	if err != nil {
		return
	}

	return category, err
}

// 获取列表
func (model *Category) List(search *string, limit int64, offset int64, order string, pid int64, categoryType string) (items []*Category, total int64, err error) {
	var (
		getItems []*Category
		getTotal int64
	)

	// 查询对象
	query := db.Client.
		Model(&model).
		Where("status", 1).
		Where("type", "ARTICLE")

	if search != nil {
		query.Where("title like ?", "%"+*search+"%")
	}

	if order != "" {
		query.Order(order)
	}

	if pid != -1 {
		query.Where("pid = ?", pid)
	}

	if categoryType != "" {
		query.Where("type = ?", categoryType)
	}

	// 查询总数量
	query.Count(&getTotal)

	if limit != 0 {
		query.Limit(int(limit))
	}

	if offset != 0 {
		query.Offset(int(offset))
	}

	// 查询列表
	query.Find(&getItems)

	// 获取错误信息
	err = query.Error
	if err != nil {
		return
	}

	return getItems, getTotal, err
}
