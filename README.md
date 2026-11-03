# Go 从零到实战：中文系统教材

一套面向初学者的 Go（Golang）中文课程：从安装环境、基础语法开始，逐步学习数据结构、函数、结构体、接口、错误处理、文件、JSON、HTTP、并发、测试，最后完成可运行的小项目。

> 目标不是背语法，而是能读懂代码、独立拆解问题、写出可维护的 Go 程序。

## 适合谁

- 没有系统学过 Go，希望从基础开始的人。
- 学过一点 Python、JavaScript 或其他语言，但不熟悉 Go 的人。
- 想用 Go 编写 CLI、API 客户端、文件处理工具、并发任务程序或 Web 服务的人。

不要求先掌握复杂算法。会使用电脑、终端，并愿意亲手敲代码即可。

## 学完以后能做什么

1. 看懂并编写普通 Go 程序。
2. 使用变量、条件、循环、函数、切片、映射、结构体和接口组织代码。
3. 创建 Go module，拆分 package，管理依赖。
4. 正确处理错误、文件、JSON 和 HTTP 请求。
5. 使用 goroutine、channel、context 和同步原语编写并发程序。
6. 编写单元测试，使用 gofmt、go vet 和 race detector 检查代码。
7. 独立完成命令行工具、日志分析器、API 客户端和简单 Web 服务。

## 学习方法

每一章按同一个顺序进行：

1. **先理解**：先读“它解决什么问题”，不要急着背定义。
2. **照着敲**：代码不要只看，要亲手输入。
3. **运行观察**：先预测输出，再运行验证。
4. **改一个地方**：改变量、条件或输入，观察结果如何变化。
5. **自己复述**：用一句话解释这段代码为什么这样写。
6. **完成练习**：先独立尝试，再看提示和参考答案。

建议每天 30–60 分钟。一次只攻克一个知识点；如果卡住，先缩小问题，不要同时引入多个新概念。

## 课程目录

| 阶段 | 内容 | 学习结果 |
|---|---|---|
| 0 | [学习路线与学习方法](docs/00-learning-map.md) | 知道先学什么、后学什么 |
| 1 | [安装 Go 和配置环境](docs/01-install.md) | 能在终端运行 Go |
| 2 | [第一个程序与 go 命令](docs/02-first-program.md) | 能创建、运行、构建程序 |
| 3 | [变量、常量与基本类型](docs/03-values-types.md) | 能保存和处理数据 |
| 4 | [条件判断与循环](docs/04-control-flow.md) | 能表达分支和重复执行 |
| 5 | [函数](docs/05-functions.md) | 能把大问题拆成小步骤 |
| 6 | [数组、切片与映射](docs/06-collections.md) | 能处理一组数据 |
| 7 | [字符串、结构体、指针与方法](docs/07-structs-pointers.md) | 能组织业务对象 |
| 8 | [包与模块](docs/08-packages-modules.md) | 能组织真实项目 |
| 9 | [接口与泛型](docs/09-interfaces-generics.md) | 能写可复用的代码 |
| 10 | [错误处理](docs/10-errors.md) | 能让程序正确面对失败 |
| 11 | [文件与 JSON](docs/11-files-json.md) | 能读写本地数据 |
| 12 | [HTTP 客户端与服务端](docs/12-http.md) | 能调用 API、创建 Web 服务 |
| 13 | [并发：goroutine、channel、context](docs/13-concurrency.md) | 能安全地执行并发任务 |
| 14 | [测试、调试与代码质量](docs/14-testing.md) | 能验证和维护程序 |
| 15 | [五个实战项目](docs/15-projects.md) | 把知识组合成完整程序 |
| 16 | [命令速查表与术语表](docs/16-cheatsheet.md) | 快速查找常用命令和概念 |
| 17 | [数据库与 database/sql](docs/17-database.md) | 连接数据库、查询和事务 |
| 18 | [配置、日志、构建与发布](docs/18-build-deploy.md) | 把程序整理成可交付项目 |

## 先运行第一个示例

需要先安装 Go。安装方法见 [第 1 章](docs/01-install.md)。

在仓库根目录执行：

~~~bash
go version
go run ./examples/hello
~~~

预期输出：

~~~text
你好，Go！
我已经成功运行第一个 Go 程序。
~~~

运行全部示例的编译与测试：

~~~bash
go test ./...
~~~

## 仓库结构

~~~text
.
├── README.md
├── PROGRESS.md
├── go.mod
├── docs/                 # 按章节编排的中文教材
├── examples/             # 每章配套的可运行代码
├── solutions/            # 各章练习参考答案
└── projects/             # 从小到大的实战项目说明
~~~

## 推荐学习节奏

- 第 1 周：环境、变量、条件、循环、函数。
- 第 2 周：切片、映射、结构体、指针、包和模块。
- 第 3 周：接口、错误、文件、JSON、HTTP。
- 第 4 周：并发、测试，并完成至少两个项目。
- 后续：把项目拆分成多个 package，补充测试、日志、配置和文档。

进度请记录在 [PROGRESS.md](PROGRESS.md)。每章练习的参考答案集中在 [solutions/](solutions/README.md)，建议先做题，再打开答案对照。

## 官方资料

- [Go 官方下载与安装](https://go.dev/doc/install)
- [Go 官方文档](https://go.dev/doc/)
- [A Tour of Go 交互式教程](https://go.dev/tour/)
- [Go 官方教程目录](https://go.dev/doc/tutorial/)
- [Go 标准库文档](https://pkg.go.dev/std)
- [Go 版本发布记录](https://go.dev/doc/devel/release)

本教材按 Go 1.27 系列编写。截至 2026-11-03，官方发布记录显示最新补丁版本为 Go 1.27.3（2026-10-19 发布）；后续版本可能更新，请以官方发布记录和标准库文档为准。

## 版权

本项目采用 MIT License。欢迎学习、练习、提出问题和提交改进。
