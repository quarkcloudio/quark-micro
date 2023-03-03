namespace go post

struct Post {
    1: i64 id (vt.gt = "0")
    2: i64 category_id
    3: string category_name
    4: string category_title
    5: string title
    6: string name
    7: string author
    8: string source
    9: string description
    10: list<string> cover_paths
    11: string content
    12: i64 view
    13: list<string> file_paths
    14: string created_at
    15: string updated_at
}

struct PageRequest {
    1: i64 id (vt.gt = "0")
    2: optional string name
}

struct PageResponse {
    1: i64 id (vt.gt = "0")
    2: string title
    3: string name
    4: string description
    5: list<string> cover_paths
    6: string content
    7: i64 view
    8: list<string> file_paths
    9: string created_at
    10: string updated_at
}

struct ArticleDetailRequest {
    1: i64 id (vt.gt = "0")
    2: optional string name
}

struct ArticleDetailResponse {
    1: i64 id (vt.gt = "0")
    2: i64 category_id
    3: string category_name
    4: string category_title
    5: string title
    6: string name
    7: string author
    8: string source
    9: string description
    10: list<string> cover_paths
    11: string content
    12: i64 view
    13: list<string> file_paths
    14: string created_at
    15: string updated_at
}

struct ArticleListRequest {
    1: optional string search
    2: i64 page (vt.ge = "1")
    3: i64 page_size (vt.ge = "0")
    4: string order
    5: i64 category_id (vt.ge = "0")
}

struct ArticleListResponse {
    1: list<Post> items
    2: i64 total
}

// 内容服务
service PostService {

    // 获取单页
    PageResponse GetPage(1: PageRequest req)
  
    // 获取文章详情
    ArticleDetailResponse GetArticleDetail(1: ArticleDetailRequest req)

    // 获取文章列表
    ArticleListResponse GetArticleList(1: ArticleListRequest req)
}