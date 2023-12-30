package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/quarkcloudio/quark-go/v2/pkg/dal"
	"github.com/quarkcloudio/quark-micro/cmd/category/rpc/config"
	category "github.com/quarkcloudio/quark-micro/cmd/category/rpc/kitex_gen/category/categoryservice"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库配置
func dbInit() {

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
}

// 服务初始化
func serviceInit() (opts []server.Option) {

	// 配置服务地址
	addr, err := net.ResolveTCPAddr("tcp", config.Service.Host)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	r, err := etcd.NewEtcdRegistry([]string{config.Registry.Host})
	if err != nil {
		log.Fatal(err)
	}

	// 配置注册中心
	opts = append(opts, server.WithRegistry(r))

	// 配置服务信息
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: config.Service.Name,
	}))

	// RPC传送元数据
	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))

	return
}

func main() {

	// 初始化数据库
	dbInit()

	// 初始化服务
	opts := serviceInit()
	svr := category.NewServer(new(CategoryServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
