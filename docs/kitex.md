kitex工具

```bash

 # 第一步、生成RPC调用代码，在cmd/post项目目录运行
kitex -module github.com/quarkcms/quark-micro ./idl/post.thrift

# 第二步、生成服务端代码，在cmd/post/rpc项目应用目录运行
kitex -module github.com/quarkcms/quark-micro -service post -use github.com/quarkcms/quark-micro/cmd/post/kitex_gen ../idl/post.thrift
```