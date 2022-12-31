package resources

import "github.com/quarkcms/quark-hertz/pkg/resource"

type Admin struct {
}

// 初始化
func (p *Admin) Init(request *resource.Request) interface{} {

	return p
}

// 字段
func (p *Admin) Fields(request *resource.Request) []interface{} {

	return []interface{}{}
}

// 搜索
func (p *Admin) Searches(request *resource.Request) []interface{} {

	return []interface{}{}
}

// 行为
func (p *Admin) Actions(request *resource.Request) []interface{} {

	return []interface{}{}
}
