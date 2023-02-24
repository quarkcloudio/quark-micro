package main

import (
	"context"
	post "github.com/quarkcms/quark-micro/cmd/post/kitex_gen/post"
)

// PostServiceImpl implements the last service interface defined in the IDL.
type PostServiceImpl struct{}

// GetPage implements the PostServiceImpl interface.
func (s *PostServiceImpl) GetPage(ctx context.Context, req *post.PageReq) (resp *post.PageResp, err error) {
	// TODO: Your code here...
	return
}

// GetArticleDetail implements the PostServiceImpl interface.
func (s *PostServiceImpl) GetArticleDetail(ctx context.Context, req *post.ArticleDetailReq) (resp *post.ArticleDetailResp, err error) {
	// TODO: Your code here...
	return
}

// GetArticleList implements the PostServiceImpl interface.
func (s *PostServiceImpl) GetArticleList(ctx context.Context, req *post.ArticleListReq) (resp *post.ArticleListResp, err error) {
	// TODO: Your code here...
	return
}
