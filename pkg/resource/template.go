package resource

type Template struct {
	TemplateType string
}

// 标记为后台登录模板
func (p *Template) IsAdminLoginTemplate() *Template {
	p.TemplateType = "adminLoginTemplate"

	return p
}
