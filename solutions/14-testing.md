# 第 14 章参考答案：测试

## 练习 14.1 和 14.2

~~~go
func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
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

## 练习 14.3

在仓库根目录运行：

~~~bash
go test ./...
go vet ./...
~~~

如果测试通过，Go 会显示各 package 的测试结果；没有测试文件的 package 会显示 [no test files]。

## 练习 14.4

越界示例：

~~~go
numbers := []int{1, 2}
fmt.Println(numbers[2])
~~~

有效索引只有 0 和 1。可以先检查 index >= 0 && index < len(numbers)，再访问切片。
