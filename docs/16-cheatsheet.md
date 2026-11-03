# 第 16 章：命令速查表与术语表

## Go 命令速查

~~~bash
go version                  # 查看版本
go env                      # 查看环境配置
go help                     # 查看命令帮助
go run .                    # 运行当前 package
go run ./examples/hello     # 运行指定示例
go build ./...              # 编译所有 package
go test ./...               # 测试所有 package
go test -v ./...            # 显示详细测试信息
go test -cover ./...        # 查看测试覆盖率
go test -race ./...         # 检查数据竞争
go vet ./...                # 静态检查
gofmt -w .                  # 格式化 Go 文件
go mod init example.com/app # 初始化 module
go mod tidy                 # 整理依赖
go list ./...               # 列出 package
go doc fmt.Println          # 查看标准库文档
~~~

## 语法速查

~~~go
// 变量
var name string = "Go"
age := 20

// 常量
const maxCount = 10

// 条件
if age >= 18 {
    fmt.Println("成年")
} else {
    fmt.Println("未成年")
}

// switch
switch age {
case 18:
    fmt.Println("刚成年")
default:
    fmt.Println("其他年龄")
}

// 循环
for i := 0; i < 3; i++ {
    fmt.Println(i)
}

// 切片
items := []string{"a", "b"}
items = append(items, "c")

// map
counts := map[string]int{"go": 1}
counts["go"]++

// 函数
func add(a, b int) int {
    return a + b
}

// 结构体
type User struct {
    Name string
    Age  int
}

// 错误处理
value, err := doSomething()
if err != nil {
    return err
}
~~~

## 常见关键字

| 关键字 | 用途 |
|---|---|
| package | 声明包 |
| import | 引入包 |
| func | 声明函数或方法 |
| var | 声明变量 |
| const | 声明常量 |
| type | 声明类型 |
| struct | 定义结构体 |
| interface | 定义接口 |
| if / else | 条件判断 |
| switch / case | 多分支判断 |
| for | 循环 |
| range | 遍历集合 |
| break | 结束循环或 switch |
| continue | 跳过本轮循环 |
| return | 返回函数 |
| defer | 延迟执行清理操作 |
| go | 启动 goroutine |
| select | 等待 channel 操作 |
| chan | 声明 channel |
| map | 映射类型 |

## 常见术语

| 术语 | 简单解释 |
|---|---|
| 编译 | 把源代码转换成可执行程序 |
| package | 一组 Go 源文件组成的代码单元 |
| module | 一组 package 的集合 |
| slice | 动态长度的序列视图 |
| map | 键值对集合 |
| struct | 把多个字段组合成一个类型 |
| pointer | 保存内存地址的值 |
| method | 绑定到某个类型的函数 |
| interface | 一组方法组成的能力描述 |
| error | 表示操作失败的值 |
| goroutine | Go 的轻量级并发执行单元 |
| channel | goroutine 之间传递数据的通道 |
| context | 传递取消信号、截止时间等请求范围信息 |
| race condition | 并发访问共享状态时出现的竞争问题 |
| unit test | 验证一个小功能是否符合预期的自动化测试 |

## 官方文档入口

- [Go Documentation](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)
- [Go Standard Library](https://pkg.go.dev/std)
- [Go Blog](https://go.dev/blog/)
- [Go Release History](https://go.dev/doc/devel/release)

## 最重要的原则

不要只记住命令。每次写程序时，先明确：

1. 输入是什么？
2. 需要经过哪些步骤？
3. 输出是什么？
4. 哪些地方可能失败？
5. 如何用测试验证？
