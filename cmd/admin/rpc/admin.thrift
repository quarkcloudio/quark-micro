namespace go admin

// 类型定义
struct Request {
  1: binary method
  2: string fullPath
  3: binary host
  4: binary path
  5: binary query
  6: binary body
}

struct ResourceRequest {
  1: Request request
}

struct ResourceResponse {
  1: binary respBody
}

// 资源服务
service Resource {
    ResourceResponse resourceHandle(1: ResourceRequest req)
}