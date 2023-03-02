namespace go banner

struct Banner {
    1: i64 id (vt.gt = "0")
    2: string title
    3: i32 url_type
    4: string url
    5: string cover_path
    6: string created_at
    7: string updated_at
}

struct BannerListRequest {
    1: string position
    2: string order
}

struct BannerListResponse {
    1: list<Banner> items
    2: i64 total
}

// 广告服务
service BannerService {

    // 获取列表
    BannerListResponse GetBannerList(1: BannerListRequest req)
}