package main

import (
	"context"

	post "github.com/quarkcms/quark-micro/cmd/post/rpc/kitex_gen/post"
	"github.com/quarkcms/quark-micro/cmd/post/rpc/model"
)

// PostServiceImpl implements the last service interface defined in the IDL.
type PostServiceImpl struct{}

// GetPage implements the PostServiceImpl interface.
func (s *PostServiceImpl) GetPage(ctx context.Context, req *post.PageRequest) (resp *post.PageResponse, err error) {
	// TODO: Your code here...
	return
}

// GetArticleDetail implements the PostServiceImpl interface.
func (s *PostServiceImpl) GetArticleDetail(ctx context.Context, req *post.ArticleDetailRequest) (resp *post.ArticleDetailResponse, err error) {
	// TODO: Your code here...
	return
}

// GetArticleList implements the PostServiceImpl interface.
func (s *PostServiceImpl) GetArticleList(ctx context.Context, req *post.ArticleListRequest) (resp *post.ArticleListResponse, err error) {
	var getItems []*post.Post

	items, total, err := (&model.Post{}).List(req.Search, req.PageSize, (req.Page-1)*req.PageSize, req.Order, req.CategoryId)
	for _, v := range items {
		getItems = append(getItems, &post.Post{
			Id:         int64(v.Id),
			CategoryId: int64(v.CategoryId),
			Name:       v.Name,
			Title:      v.Title,
			Content:    v.Content,
		})
	}

	resp = &post.ArticleListResponse{
		Items: getItems,
		Total: total,
	}

	return
}
