package template

import "github.com/quarkcms/quark-hertz/pkg/resource"

// 后台登录模板
type AdminLoginTemplate struct {
	TemplateType string      // 模板类型
	Api          string      // 登录接口
	Redirect     string      // 登录后跳转地址
	Logo         interface{} // 登录页面Logo
	Title        string      // 标题
	Description  string      // 描述
}

// 初始化模板
func (p *AdminLoginTemplate) TemplateInit() *AdminLoginTemplate {
	p.TemplateType = "adminLoginTemplate"
	p.Title = "QuarkGo"
	p.Redirect = "/index?api=/api/admin/dashboard/index"
	p.Description = "信息丰富的世界里，唯一稀缺的就是人类的注意力"

	return p
}

// 图形验证码
func (p *AdminLoginTemplate) Captcha(request *resource.Request) interface{} {

	return p
}

// 执行登录方法
func (p *AdminLoginTemplate) Handle(request *resource.Request) interface{} {

	return p
}
