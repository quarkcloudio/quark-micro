package main

import (
	"context"

	category "github.com/quarkcms/quark-micro/cmd/category/rpc/kitex_gen/category"
	"github.com/quarkcms/quark-micro/cmd/category/rpc/model"
)

// CategoryServiceImpl implements the last service interface defined in the IDL.
type CategoryServiceImpl struct{}

// GetCategoryDetail implements the CategoryServiceImpl interface.
func (s *CategoryServiceImpl) GetCategoryDetail(ctx context.Context, req *category.CategoryDetailRequest) (resp *category.CategoryDetailResponse, err error) {
	item, err := (&model.Category{}).Info(req.Id, *req.Name)

	resp = &category.CategoryDetailResponse{
		Id:          int64(item.Id),
		Title:       item.Title,
		Name:        item.Name,
		Description: item.Description,
		CreatedAt:   item.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   item.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return
}

// GetCategoryList implements the CategoryServiceImpl interface.
func (s *CategoryServiceImpl) GetCategoryList(ctx context.Context, req *category.CategoryListRequest) (resp *category.CategoryListResponse, err error) {
	var getItems []*category.Category

	items, total, err := (&model.Category{}).List(req.Search, req.PageSize, (req.Page-1)*req.PageSize, req.Order, req.Pid, req.Type)
	for _, item := range items {
		getItems = append(getItems, &category.Category{
			Id:          int64(item.Id),
			Title:       item.Title,
			Name:        item.Name,
			Description: item.Description,
			CreatedAt:   item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   item.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &category.CategoryListResponse{
		Items: getItems,
		Total: total,
	}

	return
}
