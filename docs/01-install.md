# 第 1 章：安装 Go 和配置环境

## 学习目标

- 安装 Go 编译器和工具链。
- 在终端确认 Go 已经可用。
- 知道 Go、编辑器和终端分别负责什么。

## Go 是什么

Go（也叫 Golang）是一门静态类型、编译型语言。你写下源代码，Go 工具链会检查并编译它，最后运行生成的程序。

可以先把它理解成三部分：

- **Go 语言**：你写代码时使用的语法和规则。
- **Go 工具链**：负责运行、编译、测试、格式化等工作。
- **标准库**：Go 自带的常用功能，例如文件、HTTP、JSON、时间和加密。

## 安装

官方安装说明：<https://go.dev/doc/install>

### macOS

方式一：从官网下载安装包，Apple 芯片选择 macOS ARM64，Intel 芯片选择 macOS x86-64。

方式二：如果已经使用 Homebrew：

~~~bash
brew update
brew install go
~~~

安装后关闭并重新打开终端，再执行：

~~~bash
go version
~~~

### Windows

1. 打开官方安装页面。
2. 下载 Windows MSI 安装包。
3. 双击安装，按提示完成。
4. 重新打开 PowerShell。
5. 执行：

~~~powershell
go version
~~~

### Linux

优先参考官网针对发行版的安装说明。若使用官网压缩包，注意先移除旧的 /usr/local/go 目录，再解压新版本；不要把新版本直接覆盖到旧目录里。

安装完成后确认 PATH 能找到 go 命令：

~~~bash
go version
which go
~~~

## 第一次检查环境

在终端依次执行：

~~~bash
go version
go env GOROOT GOPATH GOMOD
go help
~~~

这些命令分别表示：

- go version：查看 Go 版本。
- GOROOT：Go 工具链安装位置。
- GOPATH：Go 工作区相关目录，通常不需要手动修改。
- GOMOD：当前目录正在使用的 go.mod 文件位置；如果不在模块中，可能显示空值。
- go help：查看 Go 命令帮助。

## 编辑器

可以选择：

- Visual Studio Code + Go 扩展：免费，适合初学者。
- GoLand：功能完整的付费 IDE。
- Vim 或其他编辑器：只要能编辑纯文本即可。

初学阶段不需要同时安装多个 IDE。选一个能打开文件、运行终端和显示错误的位置即可。

## 第一次运行

进入仓库根目录，执行：

~~~bash
go run ./examples/hello
~~~

预期输出：

~~~text
你好，Go！
我已经成功运行第一个 Go 程序。
~~~

## 常见问题

### 1. 提示 command not found: go

说明终端找不到 Go 可执行文件。检查是否安装成功、PATH 是否正确，以及是否重新打开终端。

### 2. go version 显示的版本和官网不一样

这通常表示机器上安装的是另一个版本。先用 which go（macOS/Linux）或 Get-Command go（PowerShell）查看实际执行路径，再决定是否升级。

### 3. 编辑器提示缺少 gopls

gopls 是 Go 语言服务器，为编辑器提供补全、跳转和诊断。按编辑器提示安装即可。它不是运行 Go 程序的必要条件。

## 本章练习

**练习 1.1**：运行 go version，并记录版本号。

**练习 1.2**：运行 go env GOROOT GOPATH，解释这两个目录分别是什么。

**练习 1.3**：运行 examples/hello，确认出现两行中文输出。

**验收标准**：你能在终端运行 go version 和 go run ./examples/hello。
