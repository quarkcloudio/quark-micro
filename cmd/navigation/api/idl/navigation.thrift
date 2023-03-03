namespace go navigationapi

struct GetNavigationListRequest {
    1: i64 limit (api.query="limit", api.vd="len($) >= 0")
    2: i64 pid (api.query="pid", api.vd="len($) >= 0")
}

struct GetNavigationListResponse {}

service ApiService {
    GetNavigationListResponse GetNavigationList(1: GetNavigationListRequest req) (api.get="/api/v1/navigation/list")
}