package resource

import (
	"errors"
	"reflect"
	"time"

	"github.com/quarkcms/quark-hertz/pkg/component/admin/login"
)

const ADMIN_LOGIN_INDEX_ROUTE = "/api/admin/login/:resource/index"          // 后台登录渲染路由
const ADMIN_LOGIN_HANDLE_ROUTE = "/api/admin/login/:resource/handle"        // 后台登录执行路由
const ADMIN_LOGIN_CAPTCHA_ID_ROUTE = "/api/admin/login/:resource/captchaId" // 后台登录获取验证码ID路由
const ADMIN_LOGIN_CAPTCHA_ROUTE = "/api/admin/login/:resource/captcha/:id"  // 后台登录验证码路由

type AdminRoute struct {
	LoginIndexRoute     string // 后台登录渲染路由，默认值"/api/admin/login/:resource/index"
	LoginHandleRoute    string // 后台登录执行路由，默认值"/api/admin/login/:resource/handle"
	LoginCaptchaIdRoute string // 后台登录获取验证码ID路由，默认值"/api/admin/login/:resource/captchaId"
	LoginCaptchaRoute   string // 后台登录验证码路由，默认值"/api/admin/login/:resource/captcha/:id"
}

// 获取后台路由设置
func getAdminRouteConfig(adminRoute *AdminRoute) *AdminRoute {
	defaultAdminRoute := &AdminRoute{
		LoginIndexRoute:     ADMIN_LOGIN_INDEX_ROUTE,
		LoginHandleRoute:    ADMIN_LOGIN_HANDLE_ROUTE,
		LoginCaptchaIdRoute: ADMIN_LOGIN_CAPTCHA_ID_ROUTE,
		LoginCaptchaRoute:   ADMIN_LOGIN_CAPTCHA_ROUTE,
	}

	// 设置后台路由
	if adminRoute != nil {
		if adminRoute.LoginIndexRoute != "" {
			defaultAdminRoute.LoginIndexRoute = adminRoute.LoginIndexRoute
		}

		if adminRoute.LoginHandleRoute != "" {
			defaultAdminRoute.LoginHandleRoute = adminRoute.LoginHandleRoute
		}

		if adminRoute.LoginCaptchaIdRoute != "" {
			defaultAdminRoute.LoginCaptchaIdRoute = adminRoute.LoginCaptchaIdRoute
		}

		if adminRoute.LoginCaptchaRoute != "" {
			defaultAdminRoute.LoginCaptchaRoute = adminRoute.LoginCaptchaRoute
		}
	}

	return defaultAdminRoute
}

// 后台登录页面渲染
func (p *Resource) AdminLoginRender(request *Request) (interface{}, error) {
	resourceInstance := p.ResourceInstance
	if resourceInstance == nil {
		return nil, errors.New("未获取到实例")
	}

	// 默认登录接口
	defaultLoginApi := p.RouteToResourceUrl(p.AdminRoute.LoginHandleRoute)

	// 登录接口
	loginApi := reflect.
		ValueOf(resourceInstance).
		Elem().
		FieldByName("Api").String()
	if loginApi != "" {
		defaultLoginApi = loginApi
	}

	// 登录后跳转地址
	redirect := reflect.
		ValueOf(resourceInstance).
		Elem().
		FieldByName("Redirect").String()

	// Logo
	logo := reflect.
		ValueOf(resourceInstance).
		Elem().
		FieldByName("Logo").Interface()

	// 标题
	title := reflect.
		ValueOf(resourceInstance).
		Elem().
		FieldByName("Title").String()

	// 描述
	description := reflect.
		ValueOf(resourceInstance).
		Elem().
		FieldByName("Description").String()

	// 获取验证码ID链接
	captchaIdUrl := p.RouteToResourceUrl(p.AdminRoute.LoginCaptchaIdRoute)

	// 验证码链接
	captchaUrl := p.RouteToResourceUrl(p.AdminRoute.LoginCaptchaRoute)

	component := (&login.Component{}).
		SetApi(defaultLoginApi).
		SetRedirect(redirect).
		SetLogo(logo).
		SetTitle(title).
		SetDescription(description).
		SetCaptchaIdUrl(captchaIdUrl).
		SetCaptchaUrl(captchaUrl).
		SetCopyright(time.Now().Format("2006") + " QuarkGo").
		SetLinks([]map[string]interface{}{
			{
				"title": "Quark",
				"href":  "http://www.quarkcms.com/",
			},
			{
				"title": "爱小圈",
				"href":  "http://www.ixiaoquan.com",
			},
			{
				"title": "Github",
				"href":  "https://github.com/quarkcms",
			},
		}).
		JsonSerialize()

	return component, nil
}

// 返回验证码ID
func (p *Resource) AdminLoginCaptchaId(request *Request) (interface{}, error) {
	resourceInstance := p.ResourceInstance
	if resourceInstance == nil {
		return nil, errors.New("未获取到实例")
	}

	return resourceInstance.(interface {
		CaptchaId(request *Request) (interface{}, error)
	}).CaptchaId(request)
}

// 生成验证码
func (p *Resource) AdminLoginCaptcha(request *Request) (interface{}, error) {
	resourceInstance := p.ResourceInstance
	if resourceInstance == nil {
		return nil, errors.New("未获取到实例")
	}

	return resourceInstance.(interface {
		Captcha(request *Request) (interface{}, error)
	}).Captcha(request)
}

// 执行登录
func (p *Resource) AdminLoginHandle(request *Request) (interface{}, error) {
	resourceInstance := p.ResourceInstance
	if resourceInstance == nil {
		return nil, errors.New("未获取到实例")
	}

	return resourceInstance.(interface {
		Handle(request *Request) (interface{}, error)
	}).Handle(request)
}
