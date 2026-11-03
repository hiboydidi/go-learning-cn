# 第 14 章：测试、调试与代码质量

## 学习目标

- 编写 Go 单元测试。
- 理解表驱动测试。
- 使用 go test、go vet 和 race detector。
- 知道如何定位常见问题。

## 一、测试文件

Go 测试文件以 _test.go 结尾。

例如：

~~~text
calculator.go
calculator_test.go
~~~

测试函数通常以 Test 开头，参数类型是 *testing.T。

~~~go
func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5

    if got != want {
        t.Fatalf("Add(2, 3) = %d, want %d", got, want)
    }
}
~~~

## 二、表驱动测试

当一个函数需要测试多组输入时，可以用切片保存测试用例：

~~~go
func TestAddCases(t *testing.T) {
    tests := []struct {
        name string
        a    int
        b    int
        want int
    }{
        {"正数", 2, 3, 5},
        {"负数", -2, -3, -5},
        {"零", 0, 5, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Fatalf("got %d, want %d", got, tt.want)
            }
        })
    }
}
~~~

## 三、常用命令

~~~bash
go test ./...
go test -v ./...
go test -cover ./...
go test -race ./...
go vet ./...
gofmt -w .
~~~

| 命令 | 用途 |
|---|---|
| go test ./... | 测试当前模块全部 package |
| go test -v ./... | 显示更详细的测试过程 |
| go test -cover ./... | 查看测试覆盖率 |
| go test -race ./... | 检查运行中的数据竞争 |
| go vet ./... | 检查一部分可疑代码 |
| gofmt -w . | 格式化 Go 源码 |

## 四、调试顺序

遇到问题时按顺序做：

1. 读错误信息的第一处关键位置。
2. 确认文件名和行号。
3. 检查变量的实际值和类型。
4. 把复杂函数缩小成最小可复现示例。
5. 添加测试，固定住问题。
6. 修复后重新运行全部测试。

## 五、常见问题

- undefined：名称未声明、拼写错误，或 package 没有正确导入。
- cannot use X as Y：类型不匹配，需要检查类型或显式转换。
- imported and not used：导入了包但没有使用。
- declared and not used：声明了局部变量但没有使用。
- index out of range：数组或切片索引超出范围。
- nil pointer dereference：对 nil 指针进行了访问。
- concurrent map writes：多个 goroutine 同时写普通 map，需要同步或改变数据所有权设计。

## 六、测试不是最后才做

推荐先定义函数输入和输出，再写最小测试，然后实现逻辑。每修复一个 bug，都尽量增加一个能复现它的测试。

## 本章练习

**练习 14.1**：给 add 函数编写一个 TestAdd。

**练习 14.2**：用表驱动测试覆盖正数、负数和 0。

**练习 14.3**：运行 go test ./... 和 go vet ./...，记录输出。

**练习 14.4**：故意制造一个切片越界问题，理解错误信息，再修复它。

**验收标准**：能独立写一个测试，并在修改代码后重新运行测试。
