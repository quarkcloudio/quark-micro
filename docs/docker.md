由于kitex代码生成工具尚未支持Windows，Windows下需要使用docker环境开发

拉取golang镜像：
```bash
docker pull golang
```

创建容器：
```bash
docker run -it --name golang -v E:/develop:/go/develop golang /bin/bash
```

进入容器：
```bash
docker exec -it golang -it golang /bin/bash
```

配置golang环境：
```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

安装代码生成工具：
```bash
go install github.com/cloudwego/hertz/cmd/hz@latest
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```