namespace go post

struct Post {
    1: i64 Id (vt.gt = "0")
    2: string Name
}

struct PageReq {
    1: i64 Id (vt.gt = "0")
    2: optional string Name
}

struct PageResp {
    1: i64 Id
}

struct ArticleDetailReq {
    1: i64 Id (vt.gt = "0")
    2: optional string Name
}

struct ArticleDetailResp {
    1: i64 Id
}

struct ArticleListReq {
    1: optional string Query
    2: i64 Offset (vt.ge = "0")
    3: i64 Limit (vt.ge = "0")
}

struct ArticleListResp {
    1: list<Post> Items
    2: i64 Total
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