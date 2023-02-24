package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	"github.com/quarkcms/quark-go/pkg/dal"
	post "github.com/quarkcms/quark-micro/cmd/post/kitex_gen/post/postservice"
	"github.com/quarkcms/quark-micro/cmd/post/rpc/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", config.App.Host)
	svr := post.NewServer(new(PostServiceImpl), server.WithServiceAddr(addr))

	// 配置信息
	var (
		dbUser     = config.Mysql.Username
		dbPassword = config.Mysql.Password
		dbHost     = config.Mysql.Host
		dbPort     = config.Mysql.Port
		dbName     = config.Mysql.Database
		dbCharset  = config.Mysql.Charset
	)

	// 数据库配置信息
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&parseTime=True&loc=Local"

	dal.InitDB(mysql.Open(dsn), &gorm.Config{})

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
