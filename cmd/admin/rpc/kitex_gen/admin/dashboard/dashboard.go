// Code generated by Kitex v0.4.4. DO NOT EDIT.

package dashboard

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	admin "github.com/quarkcms/quark-hertz/cmd/admin/rpc/kitex_gen/admin"
)

func serviceInfo() *kitex.ServiceInfo {
	return dashboardServiceInfo
}

var dashboardServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Dashboard"
	handlerType := (*admin.Dashboard)(nil)
	methods := map[string]kitex.MethodInfo{
		"dashboardHandle": kitex.NewMethodInfo(dashboardHandleHandler, newDashboardDashboardHandleArgs, newDashboardDashboardHandleResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "admin",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func dashboardHandleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*admin.DashboardDashboardHandleArgs)
	realResult := result.(*admin.DashboardDashboardHandleResult)
	success, err := handler.(admin.Dashboard).DashboardHandle(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDashboardDashboardHandleArgs() interface{} {
	return admin.NewDashboardDashboardHandleArgs()
}

func newDashboardDashboardHandleResult() interface{} {
	return admin.NewDashboardDashboardHandleResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) DashboardHandle(ctx context.Context, req *admin.DashboardRequest) (r *admin.DashboardResponse, err error) {
	var _args admin.DashboardDashboardHandleArgs
	_args.Req = req
	var _result admin.DashboardDashboardHandleResult
	if err = p.c.Call(ctx, "dashboardHandle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}