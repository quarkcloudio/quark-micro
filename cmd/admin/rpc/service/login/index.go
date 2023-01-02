package login

import (
	"github.com/quarkcms/quark-hertz/pkg/resource"
	"github.com/quarkcms/quark-hertz/pkg/resource/template"
)

type Index struct {
	template.AdminLoginTemplate
}

// 初始化
func (p *Index) Init(request *resource.Request) interface{} {

	// 初始化模板
	p.TemplateInit()

	// 登录页面Logo
	p.Logo = false

	// 登录页面标题
	p.Title = "QuarkGo"

	// 登录页面描述
	p.Description = "信息丰富的世界里，唯一稀缺的就是人类的注意力"

	// 返回登录页面
	p.Redirect = "/index?api=/api/admin/dashboard/index"

	return p
}

// 图形验证码
func (p *Index) Captcha(request *resource.Request) interface{} {

	return p
}

// 执行登录方法
func (p *Index) Handle(request *resource.Request) interface{} {

	return p
}
