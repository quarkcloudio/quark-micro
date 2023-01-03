package resource

import (
	"reflect"
	"strings"
)

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

// 替换路由中的资源参数
//
//	url := p.RouteToResourceUrl("/api/admin/login/:resource/captchaId") // url = "/api/admin/login/index/captchaId"
func (p *Resource) RouteToResourceUrl(route string) string {
	resourceName := p.Request.ResourceName()

	return strings.ReplaceAll(route, ":resource", resourceName)
}

// 根据路由判断是否为当前加载实例
func (p *Resource) IsCurrentResourceInstance(provider interface{}) bool {
	providerName := reflect.TypeOf(provider).String()
	getNames := strings.Split(providerName, ".")
	structName := getNames[len(getNames)-1]
	resourceName := p.Request.ResourceName()

	return strings.EqualFold(strings.ToLower(structName), strings.ToLower(resourceName))
}

// 设置当前实例
func (p *Resource) SetResourceInstance(templateType string) interface{} {
	var resourceInstance interface{}

	for _, provider := range p.Providers {

		// 初始化
		resourceInstance = provider.(interface {
			Init(request *Request) interface{}
		}).Init(p.Request)

		// 模板类型
		currentTemplateType := reflect.
			ValueOf(resourceInstance).
			Elem().
			FieldByName("TemplateType").String()

		if currentTemplateType == "" {
			panic("templateType can not nil")
		}

		if currentTemplateType == templateType && p.IsCurrentResourceInstance(provider) {

			// 设置实例
			p.ResourceInstance = resourceInstance
		}
	}

	return resourceInstance
}

// 处理执行
func (p *Resource) Run() interface{} {
	var component interface{}

	// 获取request参数
	currentRequest := (&Request{}).Init(p.Request)

	switch p.Request.FullPath() {
	case p.AdminRoute.LoginIndexRoute:

		// 设置当前实例
		p.SetResourceInstance("adminLoginTemplate")

		// 渲染页面
		component = p.AdminLoginRender(currentRequest)
	case p.AdminRoute.LoginCaptchaIdRoute:

		// 设置当前实例
		p.SetResourceInstance("adminLoginTemplate")

		// 返回验证码ID
		component = p.AdminLoginCaptchaId(currentRequest)
	case p.AdminRoute.LoginCaptchaRoute:

		// 设置当前实例
		p.SetResourceInstance("adminLoginTemplate")

		// 生成验证码
		component = p.AdminLoginCaptcha(currentRequest)
	case p.AdminRoute.LoginHandleRoute:

		// 设置当前实例
		p.SetResourceInstance("adminLoginTemplate")

		// 执行登录
		component = p.AdminLoginHandle(currentRequest)
	}

	return component
}
