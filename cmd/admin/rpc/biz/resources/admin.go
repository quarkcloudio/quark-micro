package resources

type Admin struct {
}

// 初始化
func (p *Admin) Init() interface{} {

	return p
}

// 字段
func (p *Admin) Fields(request map[string]interface{}) []interface{} {

	return []interface{}{}
}

// 搜索
func (p *Admin) Searches(request map[string]interface{}) []interface{} {

	return []interface{}{}
}

// 行为
func (p *Admin) Actions(request map[string]interface{}) []interface{} {

	return []interface{}{}
}
