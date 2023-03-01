namespace go pageapi

struct GetPageRequest {
    1: i64 id (api.query="id", api.vd="len($) >= 0")
    2: string name (api.query="name", api.vd="len($) >= 0")
}

struct GetPageResponse {}

service ApiService {
    GetPageResponse GetPage(1: GetPageRequest req) (api.get="/api/v1/page/index")
}