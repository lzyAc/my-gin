# MxiqiGo 项目概览

`MxiqiGo` 是一个基于 Go 语言的微服务 / 后端框架，使用 **Gin + GORM**，遵循 **DDD（领域驱动设计）** 分层架构，支持模块化扩展。

主要特点：

1. **分层清晰**

   * **transport (HTTP / gRPC)**：处理请求/响应，不涉及业务逻辑
   * **application (Service)**：处理业务逻辑，协调不同模块
   * **domain (Entity / Repository interface)**：定义核心业务实体和接口
   * **infrastructure (Repository implementation + DB / Cache / ES / MQ 等)**：负责实际存储和外部服务调用

   业务逻辑与存储逻辑完全解耦，方便测试和扩展。

2. **模块化**

   * 每个模块（如 `user`, `org`）独立开发和维护
   * 包含自己的 `entity`、`repository` 接口、`service`、`repo` 实现和 `handler`
   * `modules.go` 统一初始化模块服务，使 `main.go` 极简

3. **可扩展性**

   * Redis 缓存 → `infrastructure/cache`
   * ES 搜索 → `infrastructure/es`
   * MQ 消息 → `infrastructure/mq`
   * gRPC 服务 → `cmd/grpc`
   * HTTP 网关 → `cmd/gateway`（可聚合多个微服务或路由，支持限流、IP 黑名单、日志埋点）
   * 新模块只需在 `modules.go` 注册，无需修改 `main.go`

---

# 目录结构

```
MxiqiGo/
├── cmd/                # 启动入口
│   ├── http/           # HTTP 服务
│   │   ├── main.go     # 程序启动
│   │   └── route.go    # 路由注册（迁移到transport/http中去了）
│   └── grpc/           # gRPC 服务入口（类似于http这样去操作，增加一个proto文件夹）
├── internal/           # 核心业务
│   ├── application/    # 应用服务层
│   │   ├── user/       # 用户业务 Service
│   │   └── org/        # 组织业务 Service
│   ├── domain/         # 领域层
│   │   ├── user/entity # 用户实体
│   │   └── user/repository.go # 用户 Repository 接口
│   ├── infrastructure/ # 基础设施层
│   │   ├── db/         # MySQL 初始化
│   │   ├── modules.go  # service/repo 构建入口
│   │   └── user/       # 用户 Repository 实现
│   └── transport/      # 传输层
│       ├── http/       # HTTP Handler
│       └── grpc/       # gRPC Handler
├── pkg/                 # 公共工具包（JWT、日志、中间件）
├── docker-compose.yml   # Docker 服务编排
├── Makefile             # 项目统一管理
└── go.mod / go.sum
```

---

# 基础依赖

* **Go** ≥ 1.18（我的版本：go1.24.12 linux/amd64）
* **Docker**（使用 `docker-compose` 启动 `mysql:8.0` 镜像，版本：29.2.1）
* **Nginx**（版本：1.28.1）

常用 Go 依赖：

* `github.com/gin-gonic/gin` → HTTP 框架
* `gorm.io/gorm` → ORM 框架
* `gorm.io/driver/mysql` → MySQL 驱动

---

# 安装流程

## 配置 Go 模块代理（国内网络可用）

```bash
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on
go env | grep GOPROXY
```

输出示例：

```
GOPROXY='https://goproxy.cn,direct'
```

## 下载依赖并检查

```bash
go mod tidy
go list -m all
```

## Docker / MySQL 配置

启动 MySQL 容器：

```bash
sudo docker compose up -d mysql
docker compose logs -f mysql
```

进入数据库：

```bash
docker exec -it my_gin_mysql mysql -uroot -pgin123
```

> 数据库初始化脚本在 `internal/infrastructure/sql/` 下，可根据需要修改。

---

# 构建与运行

```bash
# 使用 Makefile 构建
make build

# 后台运行服务
make start

# 查看服务状态
make ps

# 停止服务
make stop

# 查看日志
tail -f app.log
```

---

# API 示例

**用户注册**

```bash
curl -X POST http://localhost:8080/user/register \
  -H "Content-Type: application/json" \
  -d '{"username":"liziyue","password":"123456"}'
```

**用户登录**

```bash
curl -X POST http://localhost:8080/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"liziyue","password":"123456"}'
```

**获取用户信息**

```bash
curl -X GET http://localhost:8080/user/info
```

---

# 开发与二次开发

## 新增模块流程

1. 编写 `domain/entity` + `repository` 接口
2. 编写 `application/service`
3. 编写 `infrastructure/repo` 实现
4. 编写 `transport/handler` + `route` 注册
5. 在 `modules.go` 构建 service

## 扩展第三方服务

* Redis → `infrastructure/cache`
* ES → `infrastructure/es`
* MQ → `infrastructure/mq`

## 日志与中间件

* 日志统一使用 `pkg/logger`
* Gin 中间件统一在 `route.go` 注册

## 测试

* Service 层 → 单元测试
* Handler 层 → curl / Postman 测试

