package hertzadapter

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/quarkcms/quark-hertz/pkg/msg"
	"github.com/quarkcms/quark-hertz/pkg/resource"
)

const COMPONENT_RESPONSE = "component" // 组件类型响应
const ACTION_RESPONSE = "action"       // 行为类型响应
const FILE_RESPONSE = "file"           // 文件类型响应

// 将hertz框架的RequestContext转换为资源Request适配器
func RequestAdapter(ctx *app.RequestContext) (*resource.Request, error) {
	body, err := ctx.Body()
	if err != nil {
		return nil, err
	}

	return &resource.Request{
		MethodString:   string(ctx.Method()),
		FullPathString: ctx.FullPath(),
		HostString:     string(ctx.Host()),
		PathString:     string(ctx.Path()),
		QueryString:    string(ctx.Request.QueryString()),
		BodyBuffer:     body,
	}, nil
}

// 自动适配hertz框架Handler
func ResponseAdapter(builder *resource.Resource, responseType string, ctx *app.RequestContext) {
	request, err := RequestAdapter(ctx)
	if err != nil {
		ctx.JSON(200, msg.Error(err.Error(), ""))
		return
	}

	// 注册Request
	builder.Use(request)

	// 结果
	result, err := builder.Run()
	if err != nil {
		ctx.JSON(200, msg.Error(err.Error(), ""))
		return
	}

	switch responseType {
	case "component":

		ctx.JSON(200, result)
	case "action":

		ctx.JSON(200, msg.Success("操作成功", "", result))
	case "file":

		ctx.Write(result.([]byte))
	}
}
