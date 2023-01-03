package template

import (
	"bytes"

	"github.com/dchest/captcha"
	"github.com/quarkcms/quark-hertz/pkg/msg"
	"github.com/quarkcms/quark-hertz/pkg/resource"
)

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

// 验证码ID
func (p *AdminLoginTemplate) CaptchaId(request *resource.Request) string {
	return captcha.NewLen(4)
}

// 生成验证码
func (p *AdminLoginTemplate) Captcha(request *resource.Request) []byte {
	id := request.Param("id")
	writer := bytes.Buffer{}
	captcha.WriteImage(&writer, id, 110, 38)

	return writer.Bytes()
}

// 登录方法
func (p *AdminLoginTemplate) Handle(request *resource.Request) interface{} {

	return msg.Error("请实现登录方法", "")
}

// 退出方法
func (p *AdminLoginTemplate) Logout(request *resource.Request) interface{} {

	return msg.Error("请实现退出方法", "")
}
