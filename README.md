# Aris API Tmpl

[English](#english) | [中文](#中文)

---

<a name="english"></a>
## English

### Introduction

`aris-api-tmpl` is a production-ready Go backend template for RESTful APIs. It includes authentication, OAuth2 login, PostgreSQL, Redis, object storage, OpenAPI documentation, Docker deployment, static checks, and Git hooks.

### Features

- Fiber v2 + Huma typed API routes
- JWT authentication with access and refresh tokens
- GitHub and Google OAuth2 login
- PostgreSQL with GORM
- Redis cache
- MinIO and Tencent COS object storage adapters
- OpenAPI 3.1 documentation
- Cobra command line interface
- Zap logging with Lumberjack rotation
- Middleware for CORS, compression, recovery, trace ID, request logging, rate limiting, and permission validation
- Cron task support
- Dockerfile, local full-stack Docker Compose, production/dev Compose templates
- Makefile, static checks, and pre-commit hook

### Tech Stack

- Language: Go 1.25.1
- Framework: Fiber v2 + Huma
- Database: PostgreSQL + GORM
- Cache: Redis
- Object storage: MinIO / Tencent COS
- CLI: Cobra
- Configuration: Viper
- Logging: Zap + Lumberjack
- JSON: Sonic

### Project Structure

```text
.
├── cmd                   # Cobra CLI commands
├── docker                # Dockerfile and Compose files
├── env                   # Environment templates
├── internal              # Internal implementation
│   ├── api               # Fiber and Huma setup
│   ├── common            # Common constants, enums, models, errors
│   ├── config            # Environment configuration
│   ├── cron              # Cron jobs
│   ├── dto               # Data transfer objects
│   ├── handler           # HTTP handlers
│   ├── infrastructure    # Database, cache, storage, HTTP client, pools
│   ├── jwt               # JWT helpers
│   ├── lock              # Lock helpers
│   ├── logger            # Logger setup
│   ├── middleware        # HTTP middleware
│   ├── oauth2            # OAuth2 providers
│   ├── router            # Route registration
│   ├── service           # Business services
│   └── util              # Utility functions
├── script                # Deployment scripts
├── test                  # Integration/e2e/special tests
├── main.go
└── Makefile
```

### Quick Start

#### Prerequisites

- Go 1.25.1 or higher
- Docker and Docker Compose

#### Configure environment files

```bash
cp env/api.env.template env/api.env
cp env/postgresql.env.template env/postgresql.env
cp env/redis.env.template env/redis.env
cp env/minio.env.template env/minio.env
```

Edit the generated files and replace placeholder secrets before running non-local environments.

#### Run with Docker Compose

```bash
docker compose -f docker/docker-compose.yml up -d --build
```

This starts PostgreSQL, Redis, MinIO, the migration job, and the API server at `http://localhost:8080`.

#### Run locally

```bash
go mod download
go run main.go database migrate
go run main.go server start --host localhost --port 8080
```

### API Documentation

```text
http://localhost:8080/docs
```

### Commands

```bash
make build
make build-dev
make build-debug
make lint
make test
make test-cover

go run main.go server start --host localhost --port 8080
go run main.go database migrate
go run main.go object bucket create
go run main.go lint static ./...
```

### API Endpoints

- `GET /health` - Health check
- `GET /ssehealth` - SSE health check
- `GET /docs` - API documentation in non-production environments
- `GET /api/v1/oauth2/login?platform=github` - GitHub OAuth2 login URL
- `GET /api/v1/oauth2/login?platform=google` - Google OAuth2 login URL
- `POST /api/v1/oauth2/callback` - OAuth2 callback exchange
- `POST /api/v1/token/refresh` - Refresh JWT token
- `GET /api/v1/user/current` - Get current user info
- `PATCH /api/v1/user/` - Update current user info

### Development Checks

```bash
make lint
go test -count=1 ./...
```

Install the Git hook if desired:

```bash
bash .githooks/setup.sh
```

### License

This project is licensed under the Apache License 2.0. See `LICENSE` for details.

---

<a name="中文"></a>
## 中文

### 简介

`aris-api-tmpl` 是一个生产可用的 Go 后端模板，用于构建 RESTful API。项目内置认证、OAuth2 登录、PostgreSQL、Redis、对象存储、OpenAPI 文档、Docker 部署、静态检查和 Git hook。

### 特性

- Fiber v2 + Huma 类型化 API 路由
- JWT 访问令牌和刷新令牌
- GitHub / Google OAuth2 登录
- PostgreSQL + GORM
- Redis 缓存
- MinIO / 腾讯云 COS 对象存储适配
- OpenAPI 3.1 文档
- Cobra 命令行
- Zap + Lumberjack 日志轮转
- CORS、压缩、恢复、Trace ID、请求日志、限流、权限校验中间件
- Cron 定时任务
- Dockerfile、本地全量 Docker Compose、生产/开发 Compose 模板
- Makefile、静态检查和 pre-commit hook

### 技术栈

- 语言：Go 1.25.1
- 框架：Fiber v2 + Huma
- 数据库：PostgreSQL + GORM
- 缓存：Redis
- 对象存储：MinIO / 腾讯云 COS
- CLI：Cobra
- 配置：Viper
- 日志：Zap + Lumberjack
- JSON：Sonic

### 项目结构

```text
.
├── cmd                   # Cobra 命令
├── docker                # Dockerfile 和 Compose 文件
├── env                   # 环境变量模板
├── internal              # 内部实现
│   ├── api               # Fiber 和 Huma 初始化
│   ├── common            # 公共常量、枚举、模型、错误
│   ├── config            # 环境配置
│   ├── cron              # 定时任务
│   ├── dto               # 数据传输对象
│   ├── handler           # HTTP 处理器
│   ├── infrastructure    # 数据库、缓存、存储、HTTP client、协程池
│   ├── jwt               # JWT 工具
│   ├── lock              # 锁工具
│   ├── logger            # 日志初始化
│   ├── middleware        # HTTP 中间件
│   ├── oauth2            # OAuth2 提供方
│   ├── router            # 路由注册
│   ├── service           # 业务服务
│   └── util              # 工具函数
├── script                # 部署脚本
├── test                  # 集成/E2E/专项测试
├── main.go
└── Makefile
```

### 快速开始

#### 前置要求

- Go 1.25.1 或更高版本
- Docker 和 Docker Compose

#### 配置环境变量

```bash
cp env/api.env.template env/api.env
cp env/postgresql.env.template env/postgresql.env
cp env/redis.env.template env/redis.env
cp env/minio.env.template env/minio.env
```

非本地环境运行前，请修改生成文件中的占位密钥。

#### 使用 Docker Compose 运行

```bash
docker compose -f docker/docker-compose.yml up -d --build
```

该命令会启动 PostgreSQL、Redis、MinIO、迁移任务和 API 服务，API 地址为 `http://localhost:8080`。

#### 本地运行

```bash
go mod download
go run main.go database migrate
go run main.go server start --host localhost --port 8080
```

### API 文档

```text
http://localhost:8080/docs
```

### 常用命令

```bash
make build
make build-dev
make build-debug
make lint
make test
make test-cover

go run main.go server start --host localhost --port 8080
go run main.go database migrate
go run main.go object bucket create
go run main.go lint static ./...
```

### API 端点

- `GET /health` - 健康检查
- `GET /ssehealth` - SSE 健康检查
- `GET /docs` - 非生产环境 API 文档
- `GET /api/v1/oauth2/login?platform=github` - GitHub OAuth2 登录 URL
- `GET /api/v1/oauth2/login?platform=google` - Google OAuth2 登录 URL
- `POST /api/v1/oauth2/callback` - OAuth2 回调换取 Token
- `POST /api/v1/token/refresh` - 刷新 JWT
- `GET /api/v1/user/current` - 获取当前用户信息
- `PATCH /api/v1/user/` - 更新当前用户信息

### 开发检查

```bash
make lint
go test -count=1 ./...
```

如需安装 Git hook：

```bash
bash .githooks/setup.sh
```

### 许可证

本项目采用 Apache License 2.0 许可证，详见 `LICENSE`。
