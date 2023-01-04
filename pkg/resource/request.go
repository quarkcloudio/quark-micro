package resource

import (
	"encoding/json"
	"strings"

	"github.com/gobeam/stringy"
)

type Request struct {
	MethodString   string            `thrift:"methodString,1" frugal:"1,default,string" json:"methodString"`     // 请求方法
	FullPathString string            `thrift:"fullPathString,2" frugal:"2,default,string" json:"fullPathString"` // 路由
	HostString     string            `thrift:"hostString,3" frugal:"3,default,string" json:"hostString"`         // 主机地址
	PathString     string            `thrift:"pathString,4" frugal:"4,default,string" json:"pathString"`         // URL路径
	QueryString    string            `thrift:"queryString,5" frugal:"5,default,string" json:"queryString"`       // 请求参数
	Params         map[string]string `thrift:"params,6" frugal:"6,default,map<string:string>" json:"params"`     // URL param
	BodyBuffer     []byte            `thrift:"bodyBuffer,7" frugal:"7,default,binary" json:"bodyBuffer"`         // 请求的Body数据
}

type ParamValue struct {
	Key   int
	Value string
}

// 初始化
func (p *Request) Init(request *Request) *Request {
	p = request
	p.parseParams()

	return p
}

// Method return request method.
//
// Returned value is valid until returning from RequestHandler.
func (p *Request) Method() string {
	return p.MethodString
}

// FullPath returns a matched route full path. For not found routes
// returns an empty string.
//
//	router.GET("/user/:id", func(c *hertz.RequestContext) {
//	    c.FullPath() == "/user/:id" // true
//	})
func (p *Request) FullPath() string {
	return p.FullPathString
}

// Host returns requested host.
//
// The host is valid until returning from RequestHandler.
func (p *Request) Host() string {
	return p.HostString
}

// Path returns requested path.
//
// The path is valid until returning from RequestHandler.
func (p *Request) Path() string {
	return p.PathString
}

// Param returns the value of the URL param.
// It is a shortcut for c.Params.ByName(key)
func (p *Request) parseParams() {
	params := map[string]string{}
	fullPaths := strings.Split(p.FullPath(), "/")
	paramValues := []*ParamValue{}
	for k, v := range fullPaths {
		if strings.Contains(v, ":") {
			v = stringy.New(v).ReplaceFirst(":", "")
			mapValue := &ParamValue{
				Key:   k,
				Value: v,
			}
			paramValues = append(paramValues, mapValue)
		}
	}

	paths := strings.Split(p.Path(), "/")
	for _, v := range paramValues {
		params[v.Value] = paths[v.Key]
	}

	p.Params = params
}

// Param returns the value of the URL param.
// It is a shortcut for c.Params.ByName(key)
//
//	router.GET("/user/:id", func(c *hertz.RequestContext) {
//	    // a GET request to /user/john
//	    id := c.Param("id") // id == "john"
//	})
func (p *Request) Param(key string) string {
	return p.Params[key]
}

// ResourceName returns the value of the URL Resource param.
// If request path is "/api/admin/login/index" and route is "/api/admin/login/:resource"
//
//	resourceName := p.ResourceName() // resourceName = "index"
func (p *Request) ResourceName() string {
	return p.Param("resource")
}

// BodyParser binds the request body to a struct.
// It supports decoding the following content types based on the Content-Type header:
// application/json, application/xml, application/x-www-form-urlencoded, multipart/form-data
// If none of the content types above are matched, it will return a ErrUnprocessableEntity error
func (p *Request) BodyParser(out interface{}) error {
	return json.Unmarshal(p.BodyBuffer, out)
}
