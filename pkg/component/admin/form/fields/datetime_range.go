package fields

import "github.com/quarkcms/quark-hertz/pkg/component/admin/component"

type DatetimeRange struct {
	Item
}

// 初始化
func (p *DatetimeRange) Init() *DatetimeRange {
	p.Component = "datetimeRangeField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.DefaultValue = []interface{}{nil, nil}

	return p
}
