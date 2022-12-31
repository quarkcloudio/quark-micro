package action

import "github.com/quarkcms/quark-hertz/pkg/component/admin/component"

type Component struct {
	component.Element
	Label        string      `json:"label"`
	Block        bool        `json:"block"`
	Danger       bool        `json:"danger"`
	Disabled     bool        `json:"disabled"`
	Ghost        bool        `json:"ghost"`
	Icon         string      `json:"icon"`
	Shape        string      `json:"shape"`
	Size         string      `json:"size"`
	Type         string      `json:"type"`
	ActionType   string      `json:"actionType"`
	SubmitForm   any         `json:"submitForm"`
	Href         string      `json:"href"`
	Target       string      `json:"target"`
	Modal        interface{} `json:"modal"`
	Drawer       interface{} `json:"drawer"`
	ConfirmTitle string      `json:"confirmTitle"`
	ConfirmText  string      `json:"confirmText"`
	ConfirmType  string      `json:"confirmType"`
	Api          string      `json:"api"`
	Reload       string      `json:"reload"`
	WithLoading  bool        `json:"withLoading"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "action"
	p.Size = "default"
	p.Type = "default"

	p.SetKey("action", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 设置按钮文字
func (p *Component) SetLabel(label string) *Component {
	p.Label = label

	return p
}

// 将按钮宽度调整为其父宽度的选项
func (p *Component) SetBlock(block bool) *Component {
	p.Block = block

	return p
}

// 设置危险按钮
func (p *Component) SetDanger(danger bool) *Component {
	p.Danger = danger

	return p
}

// 按钮失效状态
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 幽灵属性，使按钮背景透明
func (p *Component) SetGhost(ghost bool) *Component {
	p.Ghost = ghost

	return p
}

// 设置按钮图标
func (p *Component) SetIcon(icon string) *Component {
	p.Icon = "icon-" + icon

	return p
}

// 设置按钮形状，可选值为 circle、 round 或者不设
func (p *Component) SetShape(shape string) *Component {
	p.Shape = shape

	return p
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Component) SetType(buttonType string, danger bool) *Component {
	p.Type = buttonType
	p.Danger = danger

	return p
}

// 设置按钮大小，large | middle | small | default
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

//【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Component) SetActionType(actionType string) *Component {
	p.ActionType = actionType

	return p
}

// 当action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
func (p *Component) SetSubmitForm(formKey string) *Component {
	p.SubmitForm = formKey

	return p
}

// 点击跳转的地址，指定此属性 button 的行为和 a 链接一致
func (p *Component) SetHref(href string) *Component {
	p.Href = href

	return p
}

// 相当于 a 链接的 target 属性，href 存在时生效
func (p *Component) SetTarget(target string) *Component {
	p.Target = target

	return p
}

// 设置跳转链接
func (p *Component) SetLink(href string, target string) *Component {
	p.SetHref(href)
	p.SetTarget(target)
	p.ActionType = "link"

	return p
}

// 弹窗
func (p *Component) SetModal(callback interface{}) *Component {
	modal := (&Modal{}).Init()
	getCallback := callback.(func(modal *Modal) interface{})

	p.Modal = getCallback(modal)

	return p
}

// 抽屉
func (p *Component) SetDrawer(callback interface{}) *Component {
	drawer := (&Drawer{}).Init()
	getCallback := callback.(func(drawer *Drawer) interface{})

	p.Drawer = getCallback(drawer)

	return p
}

// 设置行为前的确认操作
func (p *Component) SetWithConfirm(title string, text string, confirmType string) *Component {
	p.ConfirmTitle = title
	p.ConfirmText = text
	p.ConfirmType = confirmType

	return p
}

//  执行行为的接口链接
func (p *Component) SetApi(api string) *Component {
	p.Api = api
	p.ActionType = "ajax"

	return p
}

//  执行成功后刷新的组件
func (p *Component) SetReload(reload string) *Component {
	p.Reload = reload

	return p
}

//  是否具有loading
func (p *Component) SetWithLoading(loading bool) *Component {
	p.WithLoading = loading

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "action"

	return p
}
