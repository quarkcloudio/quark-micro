package service

import (
	"github.com/quarkcloudio/quark-micro/cmd/admin/biz/service/dashboard"
	"github.com/quarkcloudio/quark-micro/cmd/admin/biz/service/layout"
	"github.com/quarkcloudio/quark-micro/cmd/admin/biz/service/login"
	"github.com/quarkcloudio/quark-micro/cmd/admin/biz/service/resource"
	"github.com/quarkcloudio/quark-micro/cmd/admin/biz/service/upload"
)

// 注册服务
var Provider = []interface{}{
	&login.Index{},
	&dashboard.Index{},
	&layout.Index{},
	&resource.Article{},
	&resource.Page{},
	&resource.Category{},
	&resource.Banner{},
	&resource.BannerCategory{},
	&resource.Navigation{},
	&upload.File{},
	&upload.Image{},
}
