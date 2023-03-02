namespace go bannerapi

struct GetBannerListRequest {
    1: string position (api.query="position", api.vd="len($) > 0")
}

struct GetBannerListResponse {}

service ApiService {
    GetBannerListResponse GetBannerList(1: GetBannerListRequest req) (api.get="/api/v1/banner/list")
}