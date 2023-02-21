package handler

import (
	"github.com/quarkcms/quark-micro/cmd/admin/biz/handler/dashboard"
	"github.com/quarkcms/quark-micro/cmd/admin/biz/handler/login"
	"github.com/quarkcms/quark-micro/cmd/admin/biz/handler/resource"
)

// 注册服务
var Provider = []interface{}{
	&login.Index{},
	&dashboard.Index{},
	&resource.Article{},
	&resource.Page{},
	&resource.Category{},
	&resource.Banner{},
	&resource.BannerCategory{},
	&resource.Navigation{},
}
