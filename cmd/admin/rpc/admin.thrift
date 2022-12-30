namespace go admin

// 类型定义
struct Request {
  1: binary Body
  2: binary Query
}

struct DashboardRequest {
  1: Request request
  2: string dashboard
}

struct DashboardResponse {
  1: string respBody
}

// 仪表盘
service Dashboard {
    DashboardResponse dashboardHandle(1: DashboardRequest req)
}

struct ResourceIndexRequest {
  1: Request request
  2: string resource
  3: string search
  4: i32 page
  5: i32 pageSize
  6: string sorter
  7: string filter
}

struct ResourceIndexResponse {
  1: string respBody
}

// 列表
service ResourceIndex {
    ResourceIndexResponse resourceIndexhandle(1: ResourceIndexRequest req)
}