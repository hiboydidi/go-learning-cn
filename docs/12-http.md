# 第 12 章：HTTP 客户端与服务端

## 学习目标

- 理解 HTTP 请求和响应。
- 使用 net/http 调用 API。
- 设置超时和请求上下文。
- 创建一个简单 HTTP 服务。

## 一、HTTP 请求是什么

一次典型请求包含：

- 方法：GET、POST、PUT、DELETE 等。
- URL：请求目标。
- Header：附加信息，例如 Content-Type。
- Body：请求内容，可选。

响应通常包含：

- 状态码：例如 200、400、404、500。
- Header：响应元信息。
- Body：响应数据。

## 二、发送 GET 请求

~~~go
client := &http.Client{
    Timeout: 10 * time.Second,
}

resp, err := client.Get("https://httpbin.org/get")
if err != nil {
    return err
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
if err != nil {
    return err
}

if resp.StatusCode < 200 || resp.StatusCode >= 300 {
    return fmt.Errorf("请求失败，状态码：%d", resp.StatusCode)
}

fmt.Println(string(body))
~~~

重要事项：

1. 给客户端设置超时，避免请求无限等待。
2. 使用完响应后关闭 Body。
3. 网络请求成功不代表 HTTP 状态码一定是 2xx。
4. 对响应体大小设置合理限制，避免无边界读取不可信数据。

## 三、发送 JSON POST 请求

~~~go
payload := map[string]string{
    "name": "小明",
}

data, err := json.Marshal(payload)
if err != nil {
    return err
}

req, err := http.NewRequest(
    http.MethodPost,
    "https://httpbin.org/post",
    bytes.NewReader(data),
)
if err != nil {
    return err
}

req.Header.Set("Content-Type", "application/json")

client := &http.Client{Timeout: 10 * time.Second}
resp, err := client.Do(req)
if err != nil {
    return err
}
defer resp.Body.Close()
~~~

真实项目中，API Key 等凭据应从环境变量或安全的密钥管理系统读取，不要硬编码到源码或提交到 Git。

## 四、context：控制请求生命周期

~~~go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, err := http.NewRequestWithContext(
    ctx,
    http.MethodGet,
    "https://httpbin.org/delay/1",
    nil,
)
if err != nil {
    return err
}
~~~

context 可以传递截止时间、取消信号和请求范围内的数据。不要把 context 存在长期使用的结构体字段中；通常应作为函数的第一个参数传入。

## 五、创建一个 HTTP 服务

~~~go
mux := http.NewServeMux()

mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    io.WriteString(w, "{\"status\":\"ok\"}")
})

server := &http.Server{
    Addr:              ":8080",
    Handler:           mux,
    ReadHeaderTimeout: 5 * time.Second,
    ReadTimeout:       10 * time.Second,
    WriteTimeout:      10 * time.Second,
    IdleTimeout:       60 * time.Second,
}

log.Fatal(server.ListenAndServe())
~~~

运行服务后，在另一个终端测试：

~~~bash
curl http://localhost:8080/health
~~~

预期响应：

~~~json
{"status":"ok"}
~~~

## 六、API 客户端的基本流程

~~~mermaid
flowchart TD
    A["准备请求参数"] --> B["编码 JSON"]
    B --> C["创建 Request"]
    C --> D["设置 Header 和 Context"]
    D --> E["发送请求"]
    E --> F{"网络是否成功？"}
    F -- "否" --> G["处理网络错误"]
    F -- "是" --> H{"状态码是否为 2xx？"}
    H -- "否" --> I["读取错误响应并返回错误"]
    H -- "是" --> J["解析响应 JSON"]
    J --> K["返回业务结果"]
~~~

## 本章练习

**练习 12.1**：调用一个公开测试 API，打印状态码和响应内容。

**练习 12.2**：把 GET 请求改成 POST，并发送 JSON。

**练习 12.3**：创建 /health 接口，返回 JSON 格式的健康状态。

**练习 12.4**：给客户端增加 3 秒超时，并解释超时发生时会返回什么。

**验收标准**：能写出一个带超时、检查状态码、关闭响应 Body 的 HTTP 客户端。
