package service

import (
	"github.com/quarkcms/quark-hertz/cmd/admin/rpc/service/login"
	"github.com/quarkcms/quark-hertz/cmd/admin/rpc/service/resources"
)

// 注册服务
var Providers = []interface{}{
	&login.Login{},
	&resources.Admin{},
}
