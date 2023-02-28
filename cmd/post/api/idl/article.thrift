namespace go articleapi

struct GetArticleListRequest {
    1: optional string search (api.query="search", api.vd="len($) >= 0")
    2: i64 page (api.query="page", api.vd="len($) >= 0")
    3: i64 page_size (api.query="page_size", api.vd="len($) >= 0")
}

struct GetArticleListResponse {}

struct GetArticleDetailRequest {
    1: i64 id (api.query="id", api.vd="len($) >= 0")
    2: string name (api.query="name", api.vd="len($) >= 0")
}

struct GetArticleDetailResponse {}

service ApiService {
    GetArticleListResponse GetArticleList(1: GetArticleListRequest req) (api.get="/v1/article/list")
    GetArticleDetailResponse GetArticleDetail(1: GetArticleDetailRequest req) (api.get="/v1/article/detail")
}