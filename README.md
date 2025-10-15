# GoFurry-Fiber

### 项目描述

一个基于Fiber框架的公共类库。

### 环境依赖

|     Go版本     | 数据库 | 缓存中间件 |
| :------------: | :----: | :--------: |
| GO SDK 1.23.0+ | PGSQL  |   Redis    |

### 部署步骤

1. go build 获取二进制文件 gfs
2. chmod +777 ./gfs
3. ./gfs install 将服务注册到 systemd
4. systemctl enable gfs 这里的 gfs 可在 main.go 中进行修改
5. systemctl start gfs
