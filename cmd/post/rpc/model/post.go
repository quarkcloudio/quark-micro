package model

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-micro/cmd/post/rpc/kitex_gen/post"
	"gorm.io/gorm"
)

// 文章模型
type Post struct {
	Id            int            `json:"id" gorm:"autoIncrement"`
	Adminid       int            `json:"adminid"`
	Uid           int            `json:"uid"`
	CategoryId    int            `json:"category_id"`
	Tags          string         `json:"tags" gorm:"size:200;default:null"`
	Title         string         `json:"title" gorm:"size:200;not null"`
	Name          string         `json:"name" gorm:"size:200;default:null"`
	Author        string         `json:"author" gorm:"size:200;default:null"`
	Source        string         `json:"source" gorm:"size:200;default:null"`
	Description   string         `json:"description" gorm:"size:200;default:null"`
	Password      string         `json:"password" gorm:"size:200;default:null"`
	CoverIds      string         `json:"cover_ids" gorm:"size:1000;default:null"`
	Pid           int            `json:"pid" gorm:"default:0"`
	Level         int            `json:"level" gorm:"size:11;default:0"`
	Type          string         `json:"type" gorm:"size:200;not null;default:ARTICLE"`
	ShowType      int            `json:"show_type" gorm:"size:4;default:0"`
	Position      string         `json:"position" gorm:"size:100;default:null"`
	Link          string         `json:"link" gorm:"size:100;default:null"`
	Content       string         `json:"content" gorm:"size:5000;default:null"`
	Comment       int            `json:"comment" gorm:"default:0"`
	View          int            `json:"view" gorm:"default:0"`
	PageTpl       string         `json:"page_tpl" gorm:"size:100"`
	CommentStatus int            `json:"comment_status" gorm:"size:1;not null;default:0"`
	FileIds       string         `json:"file_ids" gorm:"size:1000;default:null"`
	Status        int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

// 获取菜单的有序列表
func (model *Post) List() (list []*post.ArticleListResp, Error error) {
	var getList []*Post

	err := db.Client.
		Model(&model).
		Where("type", "PAGE").
		Order("id asc").
		Find(&getList).Error
	if err != nil {
		return list, err
	}

	return list, nil
}
