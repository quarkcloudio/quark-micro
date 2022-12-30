package main

import (
	"context"

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
	resource := req.Resource

	return &admin.ResourceIndexResponse{
		JsonString: resource,
	}, nil
}
