package config

import (
	"github.com/quarkcms/quark-micro/pkg/env"
)

type ServiceConfig struct {
	Name string // 服务名称
	Host string // 服务地址
	Key  string // 令牌加密key，如果设置绝对不可泄漏
}

// 服务配置信息
var Service = &ServiceConfig{

	// 服务名称
	Name: env.Get("APP_NAME", "Banner"),

	// 服务地址
	Host: env.Get("APP_HOST", "127.0.0.1:8006"),

	// 令牌加密key，如果设置绝对不可泄漏
	Key: env.Get("APP_KEY", "Your App Key"),
}
