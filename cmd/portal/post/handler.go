package main

import (
	"context"
	"fmt"

	api "github.com/quarkcms/quark-micro/cmd/portal/post/kitex_gen/api"
)

// PostImpl implements the last service interface defined in the IDL.
type PostImpl struct{}

// GetPostList implements the PostImpl interface.
func (s *PostImpl) GetPostList(ctx context.Context, req *api.PostRequest) (resp *api.PostResponse, err error) {
	// TODO: Your code here...
	fmt.Println(req)
	resp = &api.PostResponse{
		RespBody: "hello world",
	}

	return resp, nil
}
