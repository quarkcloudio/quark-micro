package main

import (
	"context"

	banner "github.com/quarkcms/quark-micro/cmd/banner/rpc/kitex_gen/banner"
	"github.com/quarkcms/quark-micro/cmd/banner/rpc/model"
	"github.com/quarkcms/quark-micro/pkg/utils"
)

// BannerServiceImpl implements the last service interface defined in the IDL.
type BannerServiceImpl struct{}

// GetBannerList implements the BannerServiceImpl interface.
func (s *BannerServiceImpl) GetBannerList(ctx context.Context, req *banner.BannerListRequest) (resp *banner.BannerListResponse, err error) {
	var getItems []*banner.Banner

	items, total, err := (&model.Banner{}).List(req.Position, req.Order)
	for _, item := range items {
		getItems = append(getItems, &banner.Banner{
			Id:        int64(item.Id),
			Title:     item.Title,
			UrlType:   int32(item.UrlType),
			Url:       item.Url,
			CoverPath: utils.GetPicturePath(item.CoverId),
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: item.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &banner.BannerListResponse{
		Items: getItems,
		Total: total,
	}

	return
}
