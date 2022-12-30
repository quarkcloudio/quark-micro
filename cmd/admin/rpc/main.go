package main

import (
	admin "github.com/quarkcms/quark-hertz/cmd/admin/rpc/kitex_gen/admin/combineservice"
	"log"
)

func main() {
	svr := admin.NewServer(new(CombineServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
