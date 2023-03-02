namespace go navigation

struct Navigation {
    1: i64 id (vt.gt = "0")
    2: i64 pid
    3: string title
    4: i32 url_type
    5: string url
    6: string cover_path
    7: string created_at
    8: string updated_at
}

struct NavigationListRequest {
    1: i64 limit (vt.ge = "0")
    2: string order
    3: i64 pid (vt.ge = "0")
}

struct NavigationListResponse {
    1: list<Navigation> items
}

// 导航服务
service NavigationService {

    // 获取列表
    NavigationListResponse GetNavigationList(1: NavigationListRequest req)
}