package resource

import (
	"reflect"
)

type Request struct {
	Method   []byte `thrift:"method,1" frugal:"1,default,binary" json:"method"`     // 请求方法
	FullPath string `thrift:"fullPath,2" frugal:"2,default,string" json:"fullPath"` // 路由
	Host     []byte `thrift:"host,3" frugal:"3,default,binary" json:"host"`         // 主机地址
	Path     []byte `thrift:"path,4" frugal:"4,default,binary" json:"path"`         // URL路径
	Query    []byte `thrift:"query,5" frugal:"5,default,binary" json:"query"`       // 请求参数
	Body     []byte `thrift:"body,6" frugal:"6,default,binary" json:"body"`         // 请求的Body数据
}

const LOGIN_GET_ROUTE = "/api/admin/login"  // 后台登录渲染路由
const LOGIN_POST_ROUTE = "/api/admin/login" // 后台登录执行路由

type AdminRoute struct {
	LoginGetRoute  string // 后台登录渲染路由，默认值"/api/admin/login"
	LoginPostRoute string // 后台登录执行路由，默认值"/api/admin/login"
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
	Config           *Config
	ResourceInstance interface{}
}

// 初始化对象
func New(config *Config) *Resource {
	defaultConfig := &Config{
		Providers:  config.Providers,
		Request:    config.Request,
		AdminRoute: getAdminRouteConfig(config.AdminRoute),
	}

	return &Resource{
		Config: defaultConfig,
	}
}

// 通用配置应用方法
func (p *Resource) Use(args interface{}) *Resource {
	argsName := reflect.TypeOf(args).String()

	switch argsName {
	case "*resource.AdminRoute":
		p.Config.AdminRoute = getAdminRouteConfig(args.(*AdminRoute))
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

	return p.ResourceInstance
}

// 获取路由类型
func (p *Resource) GetRouteType() string {
	routeType := ""
	if p.Config.AdminRoute.LoginGetRoute == string(p.Config.Request.Path) {
		routeType = "adminLoginRoute"
	}

	return routeType
}

// 处理执行
func (p *Resource) Run() interface{} {
	var component interface{}
	switch p.GetRouteType() {
	case "adminLoginRoute":
		for _, provider := range p.Config.Providers {
			resourceInstance := provider.(interface {
				Init(request *Request) interface{}
			}).Init(p.Config.Request)

			// 模板类型
			templateType := reflect.
				ValueOf(resourceInstance).
				Elem().
				FieldByName("TemplateType").String()

			if templateType == "" {
				panic("templateType can not nil")
			}

			if templateType == "adminLoginTemplate" {
				p.ResourceInstance = resourceInstance
			}
		}

		component = p.AdminLoginRender()
	}

	return component
}
