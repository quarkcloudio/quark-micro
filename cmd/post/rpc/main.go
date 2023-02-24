package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	post "github.com/quarkcms/quark-micro/cmd/post/kitex_gen/post/postservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9000")
	svr := post.NewServer(new(PostServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
