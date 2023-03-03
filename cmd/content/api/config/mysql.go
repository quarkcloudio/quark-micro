package config

import (
	"github.com/quarkcms/quark-micro/pkg/env"
)

type MysqlConfig struct {
	Host     string // 地址
	Port     string // 端口
	Database string // 数据库
	Username string // 用户名
	Password string // 密码
	Charset  string // 编码
}

// Mysql配置信息
var Mysql = &MysqlConfig{

	// 地址
	Host: env.Get("DB_HOST", "127.0.0.1"),

	// 端口
	Port: env.Get("DB_PORT", "3306"),

	// 数据库
	Database: env.Get("DB_DATABASE", "quarkgo"),

	// 用户名
	Username: env.Get("DB_USERNAME", "root"),

	// 密码
	Password: env.Get("DB_PASSWORD", "root"),

	// 编码
	Charset: "utf8",
}
