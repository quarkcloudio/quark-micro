namespace go category

struct Category {
    1: i64 id (vt.gt = "0")
    2: i64 pid
    3: string title
    4: string name
    5: string description
    6: string cover_path
    7: string created_at
    8: string updated_at
}

struct CategoryDetailRequest {
    1: i64 id (vt.gt = "0")
    2: optional string name
}

struct CategoryDetailResponse {
    1: i64 id (vt.gt = "0")
    2: i64 pid
    3: string title
    4: string name
    5: string description
    6: string cover_path
    7: string created_at
    8: string updated_at
}

struct CategoryListRequest {
    1: optional string search
    2: i64 page (vt.ge = "1")
    3: i64 page_size (vt.ge = "0")
    4: string order
    5: i64 pid (vt.ge = "0")
    6: string type
}

struct CategoryListResponse {
    1: list<Category> items
    2: i64 total
}

// 通用分类服务
service CategoryService {
  
    // 获取详情
    CategoryDetailResponse GetCategoryDetail(1: CategoryDetailRequest req)

    // 获取列表
    CategoryListResponse GetCategoryList(1: CategoryListRequest req)
}