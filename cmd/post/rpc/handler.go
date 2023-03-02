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
	item, err := (&model.Post{}).Info(req.Id, *req.Name)

	resp = &post.PageResponse{
		Id:          int64(item.Id),
		Title:       item.Title,
		Name:        item.Name,
		Description: item.Description,
		Content:     item.Content,
		View:        int64(item.View),
		CreatedAt:   item.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   item.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return
}

// GetArticleDetail implements the PostServiceImpl interface.
func (s *PostServiceImpl) GetArticleDetail(ctx context.Context, req *post.ArticleDetailRequest) (resp *post.ArticleDetailResponse, err error) {
	item, err := (&model.Post{}).Info(req.Id, *req.Name)

	resp = &post.ArticleDetailResponse{
		Id:           int64(item.Id),
		CategoryId:   int64(item.CategoryId),
		CategoryName: item.CategoryName,
		Title:        item.Title,
		Name:         item.Name,
		Author:       item.Author,
		Source:       item.Source,
		Description:  item.Description,
		Content:      item.Content,
		View:         int64(item.View),
		CreatedAt:    item.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    item.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return
}

// GetArticleList implements the PostServiceImpl interface.
func (s *PostServiceImpl) GetArticleList(ctx context.Context, req *post.ArticleListRequest) (resp *post.ArticleListResponse, err error) {
	var getItems []*post.Post

	items, total, err := (&model.Post{}).List(req.Search, req.PageSize, (req.Page-1)*req.PageSize, req.Order, req.CategoryId)
	for _, v := range items {
		getItems = append(getItems, &post.Post{
			Id:           int64(v.Id),
			CategoryId:   int64(v.CategoryId),
			CategoryName: v.CategoryName,
			Title:        v.Title,
			Name:         v.Name,
			Author:       v.Author,
			Source:       v.Source,
			Description:  v.Description,
			Content:      v.Content,
			View:         int64(v.View),
			CreatedAt:    v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &post.ArticleListResponse{
		Items: getItems,
		Total: total,
	}

	return
}
