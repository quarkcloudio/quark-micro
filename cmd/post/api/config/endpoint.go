package config

type EndpointConfig struct {
	ServiceName string
	Method      string
	Tags        map[string]string
}

// 服务配置信息
var Endpoint = &EndpointConfig{

	// 服务名称
	ServiceName: "post",

	// 方法
	Method: "",

	// 标签
	Tags: nil,
}
