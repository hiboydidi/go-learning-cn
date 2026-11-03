# 示例代码

所有示例都使用 Go 标准库，不需要额外安装第三方依赖。

在仓库根目录运行：

~~~bash
go run ./examples/hello
go run ./examples/control-flow
go run ./examples/collections
go run ./examples/functions
go run ./examples/structs
go run ./examples/errors
go run ./examples/files
go run ./examples/json
go run ./examples/concurrency
go run ./examples/worker-pool
go run ./examples/logging
~~~

HTTP 示例：

~~~bash
go run ./examples/http-client
go run ./examples/http-server
~~~

http-client 会访问 httpbin.org，需要网络连接。http-server 启动后，在浏览器打开 http://localhost:8080/health 或 http://localhost:8080/hello。

运行全部示例的编译检查和测试：

~~~bash
go test ./...
~~~

建议每次只运行当前正在学习的那一个示例，先预测输出，再修改一行代码观察结果。
