项目概览

my-gin 是一个 Go 语言开发的微服务 / 后端框架，基于 Gin + GORM，遵循 DDD（领域驱动设计）分层架构，支持模块化扩展。  
1.分层清晰
  每一层级只做自己的事情
  transport(http/grpc): 处理请求/响应，不被处理业务逻辑
  application(service): 处理业务逻辑，协调不同模块
  domain(entity/repository interface): 定义核心业务实体和接口
  infrastructure(repository implementation + DB/Cache/ES/MQ...): 负责实际的存储，外部服务调用
这样业务逻辑和存储逻辑完全解耦，方便测试和扩展
2.模块化
  每个模块（如user, org）独立开发，独立维护
    有自己的entity，repository接口，service， repo 实现和handler
    modoules.go 统一初始化模块服务，让main.go 极简
3.可扩展性
  Redis 缓存 → infrastructure/cache
  ES 搜索 → infrastructure/es
  MQ 消息 → infrastructure/mq
  gRPC 服务 → cmd/grpc
  HTTP 网关 → cmd/gateway（可聚合多个微服务或路由，限流，ip黑名单，日志埋点 都能在网关层面操作）
  任何模块可以在 modules.go 注册，main.go 不需要改动。


# 目录结构
my-gin/
├── cmd/ # 启动入口
│ ├── http/ # HTTP 服务
│ │ ├── main.go # 程序启动
│ │ └── route.go # 统一注册路由
│ └── grpc/ # gRPC 服务入口
├── internal/ # 核心业务
│ ├── application/ # 应用服务层
│ │ ├── user/ # 用户业务 Service
│ │ └── org/ # 组织业务 Service
│ ├── domain/ # 领域层
│ │ ├── user/entity # 用户实体
│ │ └── user/repository.go # 用户 Repository 接口
│ ├── infrastructure/ # 基础设施层
│ │ ├── db/ # MySQL 初始化
│ │ ├── modules.go # service/repo 构建入口
│ │ └── user/ # 用户 Repository 实现
│ └── transport/ # 传输层
│ ├── http/ # HTTP Handler
│ └── grpc/ # gRPC Handler
├── pkg/ # 公共工具包（JWT、日志、中间件）
├── docker-compose.yml # Docker 服务编排
├── Makefile # 项目统一管理
└── go.mod / go.sum


# 基础依赖(我的版本)
    go => go version go1.24.12 linux/amd64
    docker（使用的docker启动mysql:8.0镜像）=>Docker version 29.2.1, build a5c7197
    nginx =>nginx version: nginx/1.28.1

# 安装流程

#配置代理
#因为国内访问/GitHub/Go 模块可能首先，需要设置Go代理
go env -w GOPROXY=https://goproxy.cn,direct
#设置 GO111MODULE 开启模块支持
go env -w GO111MODULE=on
#验证
go env | grep GOPROXY
#输出：GOPROXY='https://goproxy.cn,direct'


#确保 Go >= 1.18
go version
#下载依赖
go mod tidy
#查看依赖是否完整
go list -m all
#常用的依赖
github.com/gin-gonic/gin → HTTP 框架
gorm.io/gorm → ORM 框架
gorm.io/driver/mysql → MySQL 驱动

#Docker / Mysql 配置 项目提供docker-composer.yml 来启动mysql
#启动mysql 容器
sudo docker compose up -d mysql
#查看日志
docker compose logs -f mysql
#进入数据库
docker exec -it my_gin_mysql mysql -uroot -pgin123

#初始化数据库（docker-compose.yml 里面按照自己的需求创建就行了）

#构建项目
# 使用 Makefile 构建
make build
# 后台运行
make start
# 查看状态
make ps
# 停止服务
make stop
# 查看日志
tail -f app.log

# 用户注册
curl -X POST http://localhost:8080/user/register \
  -H "Content-Type: application/json" \
  -d '{"username":"liziyue","password":"123456"}'
# 用户登录
curl -X POST http://localhost:8080/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"liziyue","password":"123456"}'
# 获取用户信息
curl -X GET http://localhost:8080/user/info




开发与二次开发

# 新增模块流程：

写 domain/entity + repository 接口

写 application/service

写 infrastructure/repo 实现

写 transport/handler + route 注册

在 modules.go 构建 service

扩展第三方服务：

Redis → infrastructure/cache

ES → infrastructure/es

MQ → infrastructure/mq

日志 & 中间件：

日志统一使用 pkg/logger

Gin 中间件统一在 route.go 注册

测试：

Service 层单元测试

Handler 层 curl/Postman 测试
