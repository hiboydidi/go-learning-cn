# 第 1 章参考答案：环境安装

## 练习 1.1

执行：

~~~bash
go version
~~~

输出格式类似：

~~~text
go version go1.27.1 darwin/arm64
~~~

版本号和操作系统会因机器而异。

## 练习 1.2

执行：

~~~bash
go env GOROOT GOPATH
~~~

- GOROOT：Go 工具链所在目录。
- GOPATH：Go 工作区目录，常用于缓存和安装工具。

## 练习 1.3

在仓库根目录运行：

~~~bash
go run ./examples/hello
~~~

应该看到：

~~~text
你好，Go！
我已经成功运行第一个 Go 程序。
~~~

如果 go 命令找不到，检查 Go 是否安装，以及终端 PATH 是否包含 Go 可执行文件所在目录。
