namespace go post

struct Post {
    1: i64 id (vt.gt = "0")
    2: string title
    3: string name
    4: i64 category_id
    5: string content
}

struct PageReq {
    1: i64 Id (vt.gt = "0")
    2: optional string name
}

struct PageResp {
    1: i64 id
}

struct ArticleDetailReq {
    1: i64 id (vt.gt = "0")
    2: optional string name
}

struct ArticleDetailResp {
    1: i64 id
}

struct ArticleListReq {
    1: optional string query
    2: i64 offset (vt.ge = "0")
    3: i64 limit (vt.ge = "0")
}

struct ArticleListResp {
    1: list<Post> items
    2: i64 total
}

// 内容服务
service PostService {

    // 获取单页内容
    PageResp GetPage(1: PageReq req)
  
    // 获取文章详情
    ArticleDetailResp GetArticleDetail(1: ArticleDetailReq req)

    // 获取文章列表
    ArticleListResp GetArticleList(1: ArticleListReq req)
}