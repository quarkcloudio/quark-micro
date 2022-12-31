package layout

import "github.com/quarkcms/quark-hertz/pkg/component/admin/component"

type Component struct {
	component.Element
	ComponentKey  string                   `json:"componentKey"`
	Cache         bool                     `json:"cache"`
	Title         string                   `json:"title"`
	Logo          interface{}              `json:"logo"`
	Loading       bool                     `json:"loading"`
	ContentStyle  map[string]string        `json:"contentStyle"`
	HeaderActions []map[string]interface{} `json:"headerActions"`
	Layout        string                   `json:"layout"`
	HeaderTheme   string                   `json:"headerTheme"`
	SplitMenus    bool                     `json:"splitMenus"`
	ContentWidth  string                   `json:"contentWidth"`
	NavTheme      string                   `json:"navTheme"`
	PrimaryColor  string                   `json:"primaryColor"`
	FixedHeader   bool                     `json:"fixedHeader"`
	FixSiderbar   bool                     `json:"fixSiderbar"`
	IconfontUrl   string                   `json:"iconfontUrl"`
	Locale        string                   `json:"locale"`
	SiderWidth    int                      `json:"siderWidth"`
	Menu          interface{}              `json:"menu"`
	Footer        interface{}              `json:"footer"`
	Body          interface{}              `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "layout"
	p.Cache = true

	p.SetKey("layout", component.DEFAULT_CRYPT)
	p.ComponentKey = p.Key

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 是否缓存layout
func (p *Component) SetCache(cache bool) *Component {
	p.Cache = cache
	return p
}

// layout 的左上角 的 title
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// layout 的左上角的 logo
func (p *Component) SetLogo(logo interface{}) *Component {
	p.Logo = logo
	return p
}

// layout 的加载态
func (p *Component) SetLoading(loading bool) *Component {
	p.Loading = loading
	return p
}

// layout 的内容区 style
func (p *Component) SetContentStyle(contentStyle map[string]string) *Component {
	p.ContentStyle = contentStyle
	return p
}

// layout 的头部行为
func (p *Component) SetHeaderActions(headerActions []map[string]interface{}) *Component {
	p.HeaderActions = headerActions
	return p
}

// layout 的布局模式，side：右侧导航，top：顶部导航，mix：混合模式
func (p *Component) SetLayout(layout string) *Component {
	p.Layout = layout
	return p
}

// layout为mix时，顶部主题 dark | light
func (p *Component) SetHeaderTheme(headerTheme string) *Component {
	p.HeaderTheme = headerTheme
	return p
}

// layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
func (p *Component) SetContentWidth(contentWidth string) *Component {
	p.ContentWidth = contentWidth
	return p
}

// 导航的主题，'light' | 'dark'
func (p *Component) SetNavTheme(navTheme string) *Component {
	p.NavTheme = navTheme
	return p
}

// 容器控件里面的内容
func (p *Component) SetPrimaryColor(primaryColor string) *Component {
	p.PrimaryColor = primaryColor
	return p
}

// 是否固定 header 到顶部
func (p *Component) SetFixedHeader(fixedHeader bool) *Component {
	p.FixedHeader = fixedHeader
	return p
}

// 是否固定导航
func (p *Component) SetFixSiderbar(fixSiderbar bool) *Component {
	p.FixSiderbar = fixSiderbar
	return p
}

// 使用 IconFont 的图标配置
func (p *Component) SetIconfontUrl(iconfontUrl string) *Component {
	p.IconfontUrl = iconfontUrl
	return p
}

// 当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
func (p *Component) SetLocale(locale string) *Component {
	p.Locale = locale
	return p
}

// 侧边菜单宽度
func (p *Component) SetSiderWidth(siderWidth int) *Component {
	p.SiderWidth = siderWidth
	return p
}

// 自动分割菜单
func (p *Component) SetSplitMenus(splitMenus bool) *Component {
	p.SplitMenus = splitMenus
	return p
}

// 菜单
func (p *Component) SetMenu(menu interface{}) *Component {
	p.Menu = menu
	return p
}

// 页脚
func (p *Component) SetFooter(footer interface{}) *Component {
	p.Footer = footer
	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "layout"

	return p
}
