package main

import (
	"context"
	"encoding/json"

	admin "github.com/quarkcms/quark-hertz/cmd/admin/rpc/kitex_gen/admin"
	"github.com/quarkcms/quark-hertz/cmd/admin/rpc/service"
	"github.com/quarkcms/quark-hertz/pkg/resource"
)

// CombineServiceImpl implements the last service interface defined in the IDL.
type CombineServiceImpl struct{}

// ResourceHandle implements the ResourceImpl interface.
func (s *CombineServiceImpl) ResourceHandle(ctx context.Context, req *admin.ResourceRequest) (resp *admin.ResourceResponse, err error) {
	request := &resource.Request{
		Method:   req.Request.Method,
		FullPath: req.Request.FullPath,
		Host:     req.Request.Host,
		Path:     req.Request.Path,
		Query:    req.Request.Query,
		Body:     req.Request.Body,
	}

	builder := resource.New(&resource.Config{
		Request:   request,
		Providers: service.Providers,
	})

	data := builder.Run()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return &admin.ResourceResponse{}, err
	}

	return &admin.ResourceResponse{
		RespBody: jsonData,
	}, nil
}
