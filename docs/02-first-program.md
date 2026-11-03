# 第 2 章：第一个程序与 go 命令

## 学习目标

- 看懂 Go 程序的基本结构。
- 理解 package、import 和 main 函数。
- 学会运行、编译、格式化 Go 程序。

## Hello World

~~~go
package main

import "fmt"

func main() {
    fmt.Println("你好，Go！")
}
~~~

## 一行一行理解

### package main

Go 代码按 package（包）组织。可执行程序通常使用 main 包。

### import "fmt"

import 表示引入其他包。fmt 是标准库中的格式化输入输出包。

### func main()

func 用于声明函数。main 包中的 main 函数是程序的入口。

### fmt.Println(...)

调用 fmt 包中的 Println 函数，把内容打印到终端并换行。

## 程序执行顺序

~~~mermaid
flowchart TD
    A["执行 go run"] --> B["Go 工具读取源文件"]
    B --> C["检查语法和类型"]
    C --> D["编译程序"]
    D --> E["从 main 函数开始执行"]
    E --> F["调用 fmt.Println"]
    F --> G["终端显示文字"]
~~~

## 四个常用命令

### 1. 直接运行

~~~bash
go run ./examples/hello
~~~

适合学习和快速验证。

### 2. 编译

~~~bash
go build -o hello ./examples/hello
~~~

会在当前目录生成可执行文件 hello。

macOS/Linux 运行：

~~~bash
./hello
~~~

Windows PowerShell 运行：

~~~powershell
.\hello.exe
~~~

### 3. 格式化

~~~bash
gofmt -w examples/hello/main.go
~~~

Go 有统一的格式化工具。团队通常直接使用 gofmt，不需要争论缩进风格。

### 4. 查看帮助

~~~bash
go help run
go help build
go help test
~~~

## 输出更多内容

~~~go
package main

import "fmt"

func main() {
    fmt.Println("第一行")
    fmt.Println("第二行")

    fmt.Print("这段文字")
    fmt.Print("不会自动换行")

    fmt.Printf("数字：%d，文字：%s\n", 18, "Go")
}
~~~

- Println：打印后换行。
- Print：打印但不自动换行。
- Printf：按格式输出。

常见格式符：

| 格式符 | 含义 |
|---|---|
| %s | 字符串 |
| %d | 整数 |
| %f | 浮点数 |
| %t | 布尔值 |
| %v | 按默认格式输出 |
| %T | 输出变量类型 |

## 本章练习

**练习 2.1**：新建一个 main.go，打印自己的学习目标。

**练习 2.2**：用 Println 打印两行，再用 Printf 输出一个整数和一个字符串。

**练习 2.3**：把输出中的文字修改后重新运行，观察变化。

<details>
<summary>参考答案：练习 2.2</summary>

~~~go
package main

import "fmt"

func main() {
    fmt.Println("我正在学习 Go。")
    fmt.Println("今天完成第二章。")
    fmt.Printf("已经学习 %d 章，语言是 %s。\n", 2, "Go")
}
~~~

</details>

**验收标准**：你能不复制教材，独立写出包含 package、import 和 main 函数的程序。
