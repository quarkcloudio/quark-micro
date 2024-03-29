package main

import (
	"context"

	post "github.com/quarkcloudio/quark-micro/cmd/content/rpc/kitex_gen/post"
	"github.com/quarkcloudio/quark-micro/cmd/content/rpc/model"
	"github.com/quarkcloudio/quark-micro/pkg/utils"
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
		CoverPaths:  utils.GetPicturePaths(item.CoverIds),
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
		Id:            int64(item.Id),
		CategoryId:    int64(item.CategoryId),
		CategoryName:  item.CategoryName,
		CategoryTitle: item.CategoryTitle,
		Title:         item.Title,
		Name:          item.Name,
		Author:        item.Author,
		Source:        item.Source,
		Description:   item.Description,
		CoverPaths:    utils.GetPicturePaths(item.CoverIds),
		Content:       item.Content,
		View:          int64(item.View),
		CreatedAt:     item.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     item.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return
}

// GetArticleList implements the PostServiceImpl interface.
func (s *PostServiceImpl) GetArticleList(ctx context.Context, req *post.ArticleListRequest) (resp *post.ArticleListResponse, err error) {
	var getItems []*post.Post

	items, total, err := (&model.Post{}).List(req.Search, req.PageSize, (req.Page-1)*req.PageSize, req.Order, req.CategoryId)
	for _, item := range items {
		getItems = append(getItems, &post.Post{
			Id:            int64(item.Id),
			CategoryId:    int64(item.CategoryId),
			CategoryName:  item.CategoryName,
			CategoryTitle: item.CategoryTitle,
			Title:         item.Title,
			Name:          item.Name,
			Author:        item.Author,
			Source:        item.Source,
			Description:   item.Description,
			CoverPaths:    utils.GetPicturePaths(item.CoverIds),
			Content:       item.Content,
			View:          int64(item.View),
			CreatedAt:     item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:     item.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &post.ArticleListResponse{
		Items: getItems,
		Total: total,
	}

	return
}
