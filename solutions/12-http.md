# 第 12 章参考答案：HTTP

## 练习 12.1

最小 GET 请求：

~~~go
client := &http.Client{Timeout: 5 * time.Second}

resp, err := client.Get("https://httpbin.org/get")
if err != nil {
    return err
}
defer resp.Body.Close()

fmt.Println("状态码：", resp.StatusCode)

body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
if err != nil {
    return err
}
fmt.Println(string(body))
~~~

实际代码还应判断状态码是否为 2xx。

## 练习 12.2

先用 json.Marshal 编码请求体，再用 http.NewRequest 创建 POST 请求，并设置 Content-Type: application/json。具体完整流程见 examples/http-client 和第 12 章正文。

## 练习 12.3

~~~go
mux := http.NewServeMux()

mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _, _ = io.WriteString(w, "{\"status\":\"ok\"}")
})

server := &http.Server{
    Addr:              ":8080",
    Handler:           mux,
    ReadHeaderTimeout: 5 * time.Second,
}

log.Fatal(server.ListenAndServe())
~~~

## 练习 12.4

把客户端 Timeout 设置为 3 * time.Second。超过截止时间后，Do 会返回错误；如果请求使用 context，还可以通过 ctx.Err() 判断是否因为超时或取消而结束。
