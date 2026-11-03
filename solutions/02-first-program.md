# 第 2 章参考答案：第一个程序

## 练习 2.1

~~~go
package main

import "fmt"

func main() {
    fmt.Println("我的目标是学会独立编写 Go 程序。")
}
~~~

## 练习 2.2

~~~go
package main

import "fmt"

func main() {
    fmt.Println("第一行")
    fmt.Println("第二行")
    fmt.Printf("我学习了 %d 天，正在学习 %s。\n", 1, "Go")
}
~~~

## 练习 2.3

修改双引号中的文字，再运行：

~~~bash
go run .
~~~

如果程序位于独立目录，进入该目录后运行 go run .；如果使用本仓库示例，运行 go run ./examples/hello。
