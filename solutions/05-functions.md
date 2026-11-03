# 第 5 章参考答案：函数

## 练习 5.1

~~~go
func add(a, b int) int {
    return a + b
}
~~~

## 练习 5.2

~~~go
func isEven(number int) bool {
    return number%2 == 0
}
~~~

## 练习 5.3

~~~go
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
~~~

## 练习 5.4

~~~go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("除数不能为 0")
    }
    return a / b, nil
}
~~~

调用时必须先检查 err，再使用正常结果。
