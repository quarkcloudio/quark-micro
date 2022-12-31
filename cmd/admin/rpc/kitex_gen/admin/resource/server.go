// Code generated by Kitex v0.4.4. DO NOT EDIT.
package resource

import (
	server "github.com/cloudwego/kitex/server"
	admin "github.com/quarkcms/quark-hertz/cmd/admin/rpc/kitex_gen/admin"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler admin.Resource, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}