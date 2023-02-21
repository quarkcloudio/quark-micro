# QuarkHertz

Overview:

Quark-hertz is a golang micro-services based on hertz

Install:

1. 重命名.env.example 改为 .env 
2. 编辑.env文件，更改配置信息
3. 执行下面的命令完成安装：
``` bash
# 第一步，安装依赖:
go mod tidy

# 第二步，创建vendor目录
go mod vendor
```