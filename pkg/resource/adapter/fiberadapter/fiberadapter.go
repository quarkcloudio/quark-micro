package fiberadapter

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-hertz/pkg/msg"
	"github.com/quarkcms/quark-hertz/pkg/resource"
)

const COMPONENT_RESPONSE = "component" // 组件类型响应
const ACTION_RESPONSE = "action"       // 行为类型响应
const FILE_RESPONSE = "file"           // 文件类型响应

// 将gofiber框架的Ctx转换为resource框架的Request
func RequestAdapter(ctx *fiber.Ctx) (*resource.Request, error) {

	return &resource.Request{
		MethodString:   string(ctx.Method()),
		FullPathString: ctx.Route().Name,
		HostString:     string(ctx.Hostname()),
		PathString:     string(ctx.Path()),
		QueryString:    string(ctx.Context().QueryArgs().QueryString()),
		BodyBuffer:     ctx.Body(),
	}, nil
}

// 适配gofiber框架的响应
func ResponseAdapter(builder *resource.Resource, responseType string, ctx *fiber.Ctx) error {
	var responseError error
	request, err := RequestAdapter(ctx)
	if err != nil {
		return ctx.JSON(msg.Error(err.Error(), ""))
	}

	// 注册Request
	builder.Use(request)

	// 结果
	result, err := builder.Run()
	if err != nil {
		return ctx.JSON(msg.Error(err.Error(), ""))
	}

	switch responseType {
	case "component":
		return ctx.JSON(result)
	case "action":
		return ctx.JSON(msg.Success("操作成功", "", result))
	case "file":
		return ctx.SendStream(bytes.NewReader(result.([]byte)))
	}

	return responseError
}
