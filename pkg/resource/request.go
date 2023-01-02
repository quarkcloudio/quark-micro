package resource

import (
	"strings"

	"github.com/gobeam/stringy"
)

type Request struct {
	MethodString   string              `thrift:"methodString,1" frugal:"1,default,string" json:"methodString"`       // 请求方法
	FullPathString string              `thrift:"fullPathString,2" frugal:"2,default,string" json:"fullPathString"`   // 路由
	HostString     string              `thrift:"hostString,3" frugal:"3,default,string" json:"hostString"`           // 主机地址
	PathString     string              `thrift:"pathString,4" frugal:"4,default,string" json:"pathString"`           // URL路径
	QueryString    string              `thrift:"queryString,5" frugal:"5,default,string" json:"queryString"`         // 请求参数
	Params         []map[string]string `thrift:"params,6" frugal:"6,default,list<map<string:string>>" json:"params"` // URL param
	BodyBuffer     []byte              `thrift:"bodyBuffer,7" frugal:"7,default,binary" json:"bodyBuffer"`           // 请求的Body数据
}

// 初始化
func (p *Request) Init(request *Request) *Request {
	p = request

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
//
//	router.GET("/user/:id", func(c *hertz.RequestContext) {
//	    // a GET request to /user/john
//	    id := c.Param("id") // id == "john"
//	})
func (p *Request) Param(key string) string {
	fullPaths := strings.Split(p.FullPath(), ":resource")

	return stringy.New(stringy.New(string(p.Path())).ReplaceFirst(fullPaths[0], "")).ReplaceLast(fullPaths[1], "")
}
