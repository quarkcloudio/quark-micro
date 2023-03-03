package main

import (
	"context"

	models "github.com/quarkcms/quark-go/pkg/app/model"
	navigation "github.com/quarkcms/quark-micro/cmd/navigation/rpc/kitex_gen/navigation"
	"github.com/quarkcms/quark-micro/cmd/navigation/rpc/model"
)

// NavigationServiceImpl implements the last service interface defined in the IDL.
type NavigationServiceImpl struct{}

// GetNavigationList implements the NavigationServiceImpl interface.
func (s *NavigationServiceImpl) GetNavigationList(ctx context.Context, req *navigation.NavigationListRequest) (resp *navigation.NavigationListResponse, err error) {
	var getItems []*navigation.Navigation

	items, total, err := (&model.Navigation{}).List(req.Pid, req.Order)
	for _, item := range items {
		getItems = append(getItems, &navigation.Navigation{
			Id:        int64(item.Id),
			Title:     item.Title,
			UrlType:   int32(item.UrlType),
			Url:       item.Url,
			CoverPath: (&models.Picture{}).GetPath(item.CoverId),
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: item.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &navigation.NavigationListResponse{
		Items: getItems,
		Total: total,
	}

	return
}
