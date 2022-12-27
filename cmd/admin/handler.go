package main

import (
	"context"
	admin "github.com/quarkcms/quark-hertz/kitex_gen/admin"
)

// HelloImpl implements the last service interface defined in the IDL.
type HelloImpl struct{}

// Echo implements the HelloImpl interface.
func (s *HelloImpl) Echo(ctx context.Context, req *admin.Request) (resp *admin.Response, err error) {
	// TODO: Your code here...
	return
}

// Add implements the HelloImpl interface.
func (s *HelloImpl) Add(ctx context.Context, req *admin.AddRequest) (resp *admin.AddResponse, err error) {
	// TODO: Your code here...
	return
}
