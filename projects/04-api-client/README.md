# 项目 4：HTTP API 客户端

## 目标

编写一个可复用的 Go HTTP 客户端，用于调用你有权访问的公开 API 或自有服务。

## 功能要求

- 接收 API URL。
- 设置请求超时。
- 发起 GET 请求。
- 检查网络错误和 HTTP 状态码。
- 解析 JSON。
- 输出结构化结果。
- 对响应体大小设置上限。
- 测试不依赖真实线上服务。

## 推荐目录

~~~text
api-client/
├── cmd/
│   └── api-client/
│       └── main.go
├── client/
│   ├── client.go
│   └── client_test.go
└── README.md
~~~

## 推荐实现顺序

1. 定义 Client 结构体，内部保存 *http.Client。
2. 创建 NewClient，统一配置超时。
3. 编写 Get 方法，接收 context 和 URL。
4. 用 http.NewRequestWithContext 创建请求。
5. 调用 client.Do。
6. 关闭响应 Body。
7. 检查状态码。
8. 解析 JSON。
9. 使用 httptest.NewServer 编写本地测试服务器。
10. 测试正常响应、非 2xx、超时和无效 JSON。

## 错误分类建议

- 请求构造失败。
- 网络连接或超时失败。
- HTTP 状态码不是 2xx。
- 响应体读取失败。
- JSON 解码失败。

不要把所有错误都转换成一个模糊的“请求失败”。

## 安全要求

- 不要在源码中写入 API Key、Cookie、密码或访问令牌。
- 使用环境变量或安全的密钥管理工具提供凭据。
- 只调用你有权访问的 API。
- 对 URL、响应大小、超时时间和重试次数设置合理限制。
- 不要为了重试而无限循环；重试应考虑请求是否幂等。

## 验收清单

- [ ] 有统一的 HTTP Client。
- [ ] 有请求超时。
- [ ] 正确关闭 Body。
- [ ] 检查状态码。
- [ ] 能解析 JSON。
- [ ] 有本地测试服务器。
- [ ] 测试覆盖正常和错误场景。
