kitex工具

```bash

 # 生成RPC调用代码，在cmd/post项目目录运行（可选）
kitex -module github.com/quarkcloudio/quark-micro ./idl/post.thrift

# 生成服务端代码，在cmd/content/rpc项目目录运行
kitex -module github.com/quarkcloudio/quark-micro -service post ./post.thrift
```