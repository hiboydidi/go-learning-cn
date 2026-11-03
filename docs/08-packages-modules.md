# 第 8 章：包与模块

## 学习目标

- 理解 package 和 module 的关系。
- 知道如何拆分文件和目录。
- 能初始化 Go module。
- 能管理第三方依赖。

## package 和 module 有什么区别

- **package（包）**：一组放在同一目录、共同编译的 Go 源文件。
- **module（模块）**：一个或多个相关 package 的集合，由 go.mod 描述。

可以这样理解：

~~~text
一个 module
├── package A
├── package B
└── package C
~~~

## 初始化一个项目

~~~bash
mkdir myapp
cd myapp
go mod init example.com/myapp
~~~

这会生成 go.mod 文件。

如果项目会发布到 GitHub，module path 通常使用仓库地址，例如：

~~~bash
go mod init github.com/yourname/myapp
~~~

## 目录示例

~~~text
myapp/
├── go.mod
├── main.go
├── internal/
│   └── calculator/
│       └── calculator.go
└── README.md
~~~

main.go 可以调用 calculator package 中导出的函数。

## 导出规则

Go 使用名称首字母决定是否能被其他 package 使用：

- Add：首字母大写，导出。
- add：首字母小写，不导出。

示例：

~~~go
package calculator

func Add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}
~~~

其他 package 可以调用 calculator.Add，但不能调用 calculator.subtract。

## import

~~~go
import (
    "fmt"
    "github.com/yourname/myapp/internal/calculator"
)
~~~

标准库和第三方包都通过 import 引入。

## 常用模块命令

~~~bash
go mod init example.com/myapp
go mod tidy
go mod download
go list -m all
go list ./...
~~~

- go mod init：初始化模块。
- go mod tidy：整理依赖，添加缺少的依赖并移除未使用的依赖。
- go mod download：下载模块依赖。
- go list -m all：列出模块依赖。
- go list ./...：列出当前模块下的 package。

## internal 目录

Go 对 internal 目录有特殊规则：internal 下的 package 只能被规定范围内的父级目录代码导入。这适合放不希望被项目外部直接引用的实现代码。

## 什么时候拆分 package

不要为了目录看起来复杂而拆包。出现以下情况时再考虑：

- 一组功能有清晰边界。
- 多个地方需要复用同一组功能。
- 文件已经太大，难以阅读。
- 需要把内部实现和对外接口分开。

## 本章练习

**练习 8.1**：新建一个项目目录，执行 go mod init。

**练习 8.2**：创建 calculator package，写一个导出的 Add 函数。

**练习 8.3**：在 main package 中 import calculator 并调用 Add。

**验收标准**：能说清 package 和 module 的区别，并独立创建一个简单 Go module。
