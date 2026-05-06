# Engineering Capabilities Migration Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** 从 `../aris-proxy-api` 抽取可复用工程能力，迁移到 `aris-api-tmpl`，提升开发、检查、部署和 Agent 协作一致性。

**Architecture:** 不做大规模 DDD 重构，优先迁移项目级规范、脚本、静态检查、hook、Docker/Compose 和 CI 质量门禁。自定义 lint 采用轻量 CLI 包装方式，避免直接复制与 `aris-proxy-api` 业务架构强绑定的规则导致大量误报。

**Tech Stack:** Go 1.25.1、Cobra、Makefile、golangci-lint、Git hooks、Docker Buildx、Docker Compose、GitHub Actions。

---

### Task 1: Agent 指令

**Files:**
- Create: `AGENTS.md`
- Create: `CLAUDE.md`
- Create: `CODEBUDDY.md`
- Create: `.claude/settings.json`

**Steps:**
1. 以 `aris-proxy-api` 的 Agent 文档为来源。
2. 替换项目身份、命令、测试目录约束和架构描述。
3. 保持 `AGENTS.md`、`CLAUDE.md`、`CODEBUDDY.md` 内容同步。

### Task 2: 静态检查与命令行

**Files:**
- Modify: `Makefile`
- Create: `.golangci.yml`
- Create: `cmd/lint.go`
- Create: `internal/tool/lintstatic/runner.go`
- Create: `internal/tool/lintstatic/runner_test.go`

**Steps:**
1. 增加 `make lint`、`make vet`、`make fmt`、`make test-cover`。
2. 新增 `go run main.go lint static ./...`。
3. 静态检查先执行 `go vet`，如安装 `golangci-lint` 则继续执行 `golangci-lint run`。
4. 为 lint 参数归一化补单元测试。

### Task 3: Git hook 与脚本

**Files:**
- Modify: `.githooks/pre-commit`
- Modify: `.githooks/setup.sh`
- Modify: `script/deploy.sh`
- Modify: `script/deploy-dev.sh`

**Steps:**
1. pre-commit 只格式化 Git 跟踪的 Go 文件，避免扫到 worktree 或临时目录。
2. hook 运行 `go mod tidy`、`make lint`、`make test`。
3. 部署脚本进入严格模式，镜像 tag 参数化，日志默认短输出，镜像清理改为显式开关。

### Task 4: Docker / Compose / CI

**Files:**
- Modify: `docker/dockerfile`
- Create: `docker/docker-compose.yml`
- Modify: `docker/docker-compose-single.yml`
- Modify: `docker/docker-compose-dev-single.yml`
- Modify: `.github/workflows/docker-publish.yml`
- Modify: `env/*.template`

**Steps:**
1. Dockerfile 增加 BuildKit 缓存、默认 `ENTRYPOINT` 和 `CMD`。
2. 新增本地全量 Compose：PostgreSQL、Redis、MinIO、migrate、API。
3. 单服务 Compose 镜像名和 tag 参数化。
4. GitHub Actions 增加 Go 质量门禁，再做多架构镜像构建。

### Task 5: 文档与验证

**Files:**
- Modify: `README.md`
- Modify: `.gitignore`

**Steps:**
1. 修正文档中 Docker Compose 文件名和 API 路径。
2. 补 `.worktrees/`、coverage 文件忽略规则。
3. 执行 `gofmt`、`go test -count=1 ./...`、`make lint` 验证。
