# 第 18 章：配置、日志、构建与发布

## 学习目标

- 把配置和源码分开。
- 使用标准库结构化日志。
- 编译不同平台的程序。
- 知道如何整理发布产物。

## 一、配置不要写死在代码里

不推荐：

~~~go
apiKey := "真实密钥写在这里"
~~~

推荐从环境变量读取：

~~~go
apiKey := os.Getenv("SERVICE_API_KEY")
if apiKey == "" {
    return fmt.Errorf("缺少 SERVICE_API_KEY 环境变量")
}
~~~

本地可以通过终端临时设置：

~~~bash
export SERVICE_API_KEY="your-local-test-key"
~~~

不要把真实密钥写进 README、示例代码、截图或 Git 历史。已经提交过的密钥应立即轮换，仅删除文件并不能让历史中的密钥失效。

## 二、结构化日志 slog

Go 标准库提供 log/slog：

~~~go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

logger.Info(
    "任务完成",
    "task_id", 123,
    "duration_ms", 42,
)
~~~

结构化日志把事件和字段分开，方便检索和统计。

日志中不要写入密码、访问令牌、完整 Cookie 或不必要的个人数据。

## 三、构建程序

在 module 根目录：

~~~bash
go build ./...
~~~

构建一个可执行文件：

~~~bash
go build -o bin/myapp ./cmd/myapp
~~~

安装命令行程序到 Go 的 bin 目录：

~~~bash
go install ./cmd/myapp
~~~

## 四、跨平台构建

Go 可以通过 GOOS 和 GOARCH 指定目标平台：

~~~bash
GOOS=linux GOARCH=amd64 go build -o myapp-linux-amd64 ./cmd/myapp
GOOS=linux GOARCH=arm64 go build -o myapp-linux-arm64 ./cmd/myapp
GOOS=darwin GOARCH=arm64 go build -o myapp-macos-arm64 ./cmd/myapp
GOOS=windows GOARCH=amd64 go build -o myapp-windows-amd64.exe ./cmd/myapp
~~~

交叉编译时，某些依赖 cgo 的项目还需要额外的交叉编译器和系统库。纯 Go 项目通常更简单。

## 五、嵌入静态文件

标准库 embed 可以把静态文件嵌入程序：

~~~go
import "embed"

//go:embed templates/*
var templates embed.FS
~~~

使用前需要了解 embed 的路径规则。不要把密钥、用户数据或运行时配置嵌入二进制文件。

## 六、发布前检查

~~~bash
gofmt -w .
go test ./...
go test -race ./...
go vet ./...
go build ./...
~~~

建议发布时记录：

- 版本号。
- Git commit。
- Go 版本。
- 支持的平台。
- 构建命令。
- 配置项说明。
- 已知限制。

## 七、版本管理

常见语义化版本形式：

- v0.1.0：早期开发版本。
- v1.0.0：第一个稳定主版本。
- v1.1.0：增加向后兼容功能。
- v1.1.1：修复问题。

Go module 对 v2 及以上主版本有路径规则，发布前应阅读官方模块文档。

## 本章练习

**练习 18.1**：把一个配置项从硬编码改成环境变量。

**练习 18.2**：使用 slog 输出一条 JSON 格式日志。

**练习 18.3**：分别构建 Linux amd64 和 macOS arm64 版本。

**练习 18.4**：写一个发布检查清单，包含测试、格式化、构建和版本号。

**验收标准**：能把程序配置、日志、构建和发布流程与业务代码分开。
