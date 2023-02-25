package config

import (
	"github.com/quarkcms/quark-micro/pkg/env"
)

type RegistryConfig struct {
	Host     string // 注册中心服务地址
	Username string // 注册中心用户名
	Password string // 注册中心密码
}

// 注册中心配置信息
var Registry = &RegistryConfig{

	// 注册中心服务地址
	Host: env.Get("REGISTRY_HOST", "127.0.0.1:2379"),

	// 注册中心用户名
	Username: env.Get("REGISTRY_USERNAME", ""),

	// 注册中心密码
	Password: env.Get("REGISTRY_PASSWORD", ""),
}
