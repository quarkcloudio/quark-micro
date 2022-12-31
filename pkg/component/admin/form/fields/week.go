package fields

import "github.com/quarkcms/quark-hertz/pkg/component/admin/component"

type Week struct {
	Item
}

// 初始化
func (p *Week) Init() *Week {
	p.Component = "weekField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}
