# 第 18 章参考答案：构建与发布

## 练习 18.1

~~~go
apiKey := os.Getenv("SERVICE_API_KEY")
if apiKey == "" {
    return fmt.Errorf("缺少 SERVICE_API_KEY")
}
~~~

## 练习 18.2

~~~go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("任务完成", "task_id", 123, "duration_ms", 42)
~~~

## 练习 18.3

~~~bash
GOOS=linux GOARCH=amd64 go build -o app-linux-amd64 .
GOOS=darwin GOARCH=arm64 go build -o app-macos-arm64 .
~~~

交叉编译结果取决于项目依赖。使用 cgo 的项目可能需要目标平台的 C 编译器。

## 练习 18.4

发布前至少检查：

- gofmt -w .
- go test ./...
- go test -race ./...
- go vet ./...
- go build ./...
- 确认没有提交密钥和本地数据。
- 记录版本号、commit 和支持的平台。
