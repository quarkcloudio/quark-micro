package template

import (
	"bytes"
	"errors"

	"github.com/dchest/captcha"
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

// 初始化
func (p *AdminLoginTemplate) Init(request *resource.Request) interface{} {

	// 初始化模板
	p.TemplateInit()

	return p
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
func (p *AdminLoginTemplate) CaptchaId(request *resource.Request) (interface{}, error) {

	return map[string]string{
		"captchaId": captcha.NewLen(4),
	}, nil
}

// 生成验证码
func (p *AdminLoginTemplate) Captcha(request *resource.Request) (interface{}, error) {
	id := request.Param("id")
	writer := bytes.Buffer{}
	captcha.WriteImage(&writer, id, 110, 38)

	return writer.Bytes(), nil
}

// 登录方法
func (p *AdminLoginTemplate) Handle(request *resource.Request) (interface{}, error) {

	return nil, errors.New("请实现登录方法")
}

// 退出方法
func (p *AdminLoginTemplate) Logout(request *resource.Request) (interface{}, error) {

	return nil, errors.New("请实现退出方法")
}
