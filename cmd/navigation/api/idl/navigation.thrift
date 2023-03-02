namespace go categoryapi

struct GetCategoryListRequest {
    1: optional string search (api.query="search", api.vd="len($) >= 0")
    2: i64 page (api.query="page", api.vd="len($) >= 0")
    3: i64 page_size (api.query="page_size", api.vd="len($) >= 0")
    4: i64 pid (api.query="pid", api.vd="len($) >= 0")
    5: string type (api.query="type", api.vd="len($) >= 0")
}

struct GetCategoryListResponse {}

struct GetCategoryDetailRequest {
    1: i64 id (api.query="id", api.vd="len($) >= 0")
    2: string name (api.query="name", api.vd="len($) >= 0")
}

struct GetCategoryDetailResponse {}

service ApiService {
    GetCategoryListResponse GetCategoryList(1: GetCategoryListRequest req) (api.get="/api/v1/category/list")
    GetCategoryDetailResponse GetCategoryDetail(1: GetCategoryDetailRequest req) (api.get="/api/v1/category/detail")
}