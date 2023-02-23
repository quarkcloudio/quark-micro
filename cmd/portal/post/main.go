package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	api "github.com/quarkcms/quark-micro/cmd/portal/post/kitex_gen/api/post"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9000")
	svr := api.NewServer(new(PostImpl), server.WithServiceAddr(addr))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
