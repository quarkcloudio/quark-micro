package main

import (
	admin "github.com/quarkcms/quark-hertz/kitex_gen/admin/hello"
	"log"
)

func main() {
	svr := admin.NewServer(new(HelloImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
