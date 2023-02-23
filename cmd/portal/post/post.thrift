namespace go api

struct PostRequest {
  1: string query
  2: i32 page
  3: i32 pageSize
  4: string sorter
}

struct PostResponse {
  1: string respBody
}

// 文章
service Post {
  
    // 获取文章列表
    PostResponse getPostList(1: PostRequest req)
}