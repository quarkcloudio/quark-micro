package config

import (
	"github.com/quarkcms/quark-micro/pkg/env"
)

type AppConfig struct {
	Version string // 应用版本
	Name    string // 应用名称
	Env     string // 项目环境
	Host    string // 服务地址
	Key     string // 令牌加密key，如果设置绝对不可泄漏
}

// APP配置信息
var App = &AppConfig{

	// 应用版本
	Version: "1.0.0",

	// 应用名称
	Name: env.Get("APP_NAME", "QuarkMicro").(string),

	// 项目环境
	Env: env.Get("APP_ENV", "local").(string),

	// 服务地址
	Host: env.Get("APP_HOST", "127.0.0.1:3000").(string),

	// 令牌加密key，如果设置绝对不可泄漏
	Key: env.Get("APP_KEY").(string),
}
