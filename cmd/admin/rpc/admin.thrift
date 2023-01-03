namespace go admin

// 类型定义
struct Request {
  1: string methodString // 请求方法
  2: string fullPathString // 路由
  3: string hostString // 主机地址
  4: string pathString // URL路径
  5: string queryString // 请求参数
  6: map<string, string> params  // URL param
  7: binary bodyBuffer // 请求的Body数据
}

struct ResourceRequest {
  1: Request request
}

struct ResourceResponse {
  1: binary respBody
}

// 资源服务
service Resource {
    ResourceResponse resourceHandle(1: ResourceRequest req) // 通用资源服务
    ResourceResponse captchaHandle(1: ResourceRequest req) // 生成验证码服务
}