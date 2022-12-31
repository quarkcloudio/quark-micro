package page

import "github.com/quarkcms/quark-hertz/pkg/component/admin/component"

type Component struct {
	component.Element
	Title string      `json:"title"`
	Body  interface{} `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "page"

	p.SetKey("page", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "page"

	return p
}
