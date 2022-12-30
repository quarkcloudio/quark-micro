package fields

import "github.com/quarkcms/quark-hertz/pkg/admin/component"

type Quarter struct {
	Item
}

// 初始化
func (p *Quarter) Init() *Quarter {
	p.Component = "quarterField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}
