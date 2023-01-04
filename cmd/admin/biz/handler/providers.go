package handler

import (
	"github.com/quarkcms/quark-hertz/cmd/admin/biz/handler/login"
	"github.com/quarkcms/quark-hertz/cmd/admin/biz/handler/resources"
)

// 注册服务
var Providers = []interface{}{
	&login.Index{},
	&resources.Admin{},
}
