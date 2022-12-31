package login

import "github.com/quarkcms/quark-hertz/pkg/resource"

type Login struct {
	resource.Template
}

// 初始化
func (p *Login) Init(request *resource.Request) interface{} {

	// 使用后台登录模板
	p.IsAdminLoginTemplate()

	return p
}

// 标题
func (p *Login) Title(request *resource.Request) string {

	return "QuarkGo"
}
