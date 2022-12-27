namespace go admin

typedef map<string, string> Data  // 类型定义

struct DashboardRequest {
  1: string dashboard
}

struct DashboardResponse {
  1: Data dashboard
}

service Dashboard {
    DashboardResponse dashboardHandle(1: DashboardRequest req)
}

struct ResourceIndexRequest {
  1: string resource
  2: Data search
  3: i32 page
  4: i32 pageSize
  5: string sorter
  6: string filter
}

struct ResourceIndexResponse {
  1: Data resource
}

service ResourceIndex {
    ResourceIndexResponse resourceIndexhandle(1: ResourceIndexRequest req)
}