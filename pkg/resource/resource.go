package resource

import (
	"reflect"
	"strings"
	"time"

	"github.com/quarkcms/quark-hertz/pkg/component/admin/login"
)

const LOGIN_GET_ROUTE = "/api/admin/login/:resource"  // 后台登录渲染路由
const LOGIN_POST_ROUTE = "/api/admin/login/:resource" // 后台登录执行路由

type AdminRoute struct {
	LoginGetRoute  string // 后台登录渲染路由，默认值"/api/admin/login/index"
	LoginPostRoute string // 后台登录执行路由，默认值"/api/admin/login/index"
}

type Config struct {
	Providers  []interface{} // 服务列表
	Request    *Request      // 请求数据
	AdminRoute *AdminRoute   // 后台路由
}

type Resourcer interface {
	Init()
	Fields(request *Request) []interface{}
	Searches(request *Request) []interface{}
	Actions(request *Request) []interface{}
}

type Resource struct {
	Providers        []interface{} // 服务列表
	Request          *Request      // 请求数据
	AdminRoute       *AdminRoute   // 后台路由
	ResourceInstance interface{}
}

// 初始化对象
func New(config *Config) *Resource {

	return &Resource{
		Providers:  config.Providers,
		Request:    config.Request,
		AdminRoute: getAdminRouteConfig(config.AdminRoute),
	}
}

// 通用配置应用方法
func (p *Resource) Use(args interface{}) *Resource {
	argsName := reflect.TypeOf(args).String()

	switch argsName {
	case "*resource.AdminRoute":
		p.AdminRoute = getAdminRouteConfig(args.(*AdminRoute))
	}

	return p
}

// 获取后台路由设置
func getAdminRouteConfig(adminRoute *AdminRoute) *AdminRoute {
	defaultAdminRoute := &AdminRoute{
		LoginGetRoute:  LOGIN_GET_ROUTE,
		LoginPostRoute: LOGIN_POST_ROUTE,
	}

	// 设置后台路由
	if adminRoute != nil {
		if adminRoute.LoginGetRoute != "" {
			defaultAdminRoute.LoginGetRoute = adminRoute.LoginGetRoute
		}

		if adminRoute.LoginPostRoute != "" {
			defaultAdminRoute.LoginPostRoute = adminRoute.LoginPostRoute
		}
	}

	return defaultAdminRoute
}

// 后台登录页面渲染
func (p *Resource) AdminLoginRender() interface{} {
	resourceInstance := p.ResourceInstance
	if resourceInstance == nil {
		return nil
	}

	// 默认登录接口
	defaultLoginApi := p.AdminRoute.LoginPostRoute

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

	component := (&login.Component{}).
		SetApi(defaultLoginApi).
		SetRedirect(redirect).
		SetLogo(logo).
		SetTitle(title).
		SetDescription(description).
		SetCaptchaUrl("/api/admin/captcha").
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

	return component
}

// 获取路由类型
func (p *Resource) GetRouteType() string {
	routeType := ""
	if p.AdminRoute.LoginGetRoute == p.Request.FullPath() {
		routeType = "adminLoginRoute"
	}

	return routeType
}

// 根据路由判断是否为当前加载实例
func (p *Resource) IsCurrentResourceInstance(provider interface{}) bool {
	providerName := reflect.TypeOf(provider).String()
	getNames := strings.Split(providerName, ".")
	structName := getNames[len(getNames)-1]
	resourceName := p.Request.Param("resource")

	return strings.EqualFold(strings.ToLower(structName), strings.ToLower(resourceName))
}

// 处理执行
func (p *Resource) Run() interface{} {
	var component interface{}
	switch p.GetRouteType() {
	case "adminLoginRoute":
		for _, provider := range p.Providers {

			// 初始化
			resourceInstance := provider.(interface {
				Init(request *Request) interface{}
			}).Init(p.Request)

			// 模板类型
			templateType := reflect.
				ValueOf(resourceInstance).
				Elem().
				FieldByName("TemplateType").String()

			if templateType == "" {
				panic("templateType can not nil")
			}

			if templateType == "adminLoginTemplate" && p.IsCurrentResourceInstance(provider) {

				// 设置实例
				p.ResourceInstance = resourceInstance
			}
		}

		component = p.AdminLoginRender()
	}

	return component
}
