package main

import (
	"context"
	"fmt"

	admin "github.com/quarkcms/quark-hertz/cmd/admin/rpc/kitex_gen/admin"
)

// CombineServiceImpl implements the last service interface defined in the IDL.
type CombineServiceImpl struct{}

// DashboardHandle implements the DashboardImpl interface.
func (s *CombineServiceImpl) DashboardHandle(ctx context.Context, req *admin.DashboardRequest) (resp *admin.DashboardResponse, err error) {
	// TODO: Your code here...
	return
}

// ResourceIndexhandle implements the ResourceIndexImpl interface.
func (s *CombineServiceImpl) ResourceIndexhandle(ctx context.Context, req *admin.ResourceIndexRequest) (resp *admin.ResourceIndexResponse, err error) {
	query := req.Request.Query
	body := req.Request.Body
	resource := req.Resource

	fmt.Print(query)
	fmt.Print(body)

	return &admin.ResourceIndexResponse{
		RespBody: resource,
	}, nil
}
