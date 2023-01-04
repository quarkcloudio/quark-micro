package login

import (
	"errors"
	"fmt"

	"github.com/quarkcms/quark-hertz/pkg/resource"
	"github.com/quarkcms/quark-hertz/pkg/resource/dal/db"
	"github.com/quarkcms/quark-hertz/pkg/resource/model"
	"github.com/quarkcms/quark-hertz/pkg/resource/template"
)

type Index struct {
	template.AdminLoginTemplate
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Captcha  string `json:"captcha" form:"captcha"`
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

	// 登录后跳转地址
	p.Redirect = "/index?api=/api/admin/dashboard/index"

	return p
}

// 登录方法
func (p *Index) Handle(request *resource.Request) (interface{}, error) {
	loginRequest := &LoginRequest{}
	if err := request.BodyParser(loginRequest); err != nil {
		return nil, err
	}

	if loginRequest.Captcha == "" {
		return nil, errors.New("验证码不能为空！")
	}

	if loginRequest.Username == "" || loginRequest.Password == "" {
		return nil, errors.New("用户名或密码不能为空")
	}

	var admin model.Admin
	db.DB.First(&admin)
	fmt.Print(admin)

	return nil, errors.New("请实现登录方法")
}

// 退出方法
func (p *Index) Logout(request *resource.Request) (interface{}, error) {

	return "退出成功", nil
}
