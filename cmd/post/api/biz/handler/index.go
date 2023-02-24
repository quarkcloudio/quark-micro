// Code generated by hertz generator.

package handler

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/quarkcms/quark-micro/cmd/post/kitex_gen/post"
	"github.com/quarkcms/quark-micro/cmd/post/kitex_gen/post/postservice"
	"github.com/quarkcms/quark-micro/cmd/post/rpc/config"
)

// 首页
func Index(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, utils.H{
		"message": "Hello World!",
	})
}

// 获取文章列表
func List(ctx context.Context, c *app.RequestContext) {
	postClient, err := postservice.NewClient("post", client.WithHostPorts(config.App.Host))
	if err != nil {
		log.Fatal(err)
	}

	req := &post.ArticleListReq{}
	resp, err := postClient.GetArticleList(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, utils.H{
		"message": resp.Items,
	})
}
