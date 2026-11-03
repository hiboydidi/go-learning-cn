# 第 8 章参考答案：包与模块

## 练习 8.1

~~~bash
mkdir calculator-app
cd calculator-app
go mod init example.com/calculator-app
~~~

## 练习 8.2

创建 calculator/calculator.go：

~~~go
package calculator

func Add(a, b int) int {
    return a + b
}
~~~

## 练习 8.3

在 main.go 中调用：

~~~go
package main

import (
    "fmt"
    "example.com/calculator-app/calculator"
)

func main() {
    fmt.Println(calculator.Add(2, 3))
}
~~~

这里的 import 路径必须与 go.mod 中的 module path 和目录结构对应。运行 go run .，预期输出 5。
