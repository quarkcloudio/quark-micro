package config

import (
	"github.com/quarkcloudio/quark-micro/pkg/env"
)

type Endpoint struct {
	ServiceName string
	Method      string
	Tags        map[string]string
}

type RegistryConfig struct {
	Host     string    // 服务地址
	Username string    // 用户名
	Password string    // 密码
	Endpoint *Endpoint // RPC服务端点
}

// 注册中心配置
var Registry = &RegistryConfig{

	// 服务地址
	Host: env.Get("REGISTRY_HOST", "127.0.0.1:2379").(string),

	// 用户名
	Username: env.Get("REGISTRY_USERNAME", "").(string),

	// 密码
	Password: env.Get("REGISTRY_PASSWORD", "").(string),

	// RPC服务端点
	Endpoint: &Endpoint{

		// 服务名称
		ServiceName: "Navigation",

		// 方法
		Method: "",

		// 标签
		Tags: nil,
	},
}
