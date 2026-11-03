# 第 10 章参考答案：错误处理

## 练习 10.1

~~~go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为 0")
    }
    return a / b, nil
}
~~~

## 练习 10.2

~~~go
var ErrInvalidAge = errors.New("年龄不能小于 0")

func validateAge(age int) error {
    if age < 0 {
        return ErrInvalidAge
    }
    return nil
}
~~~

## 练习 10.3

~~~go
var ErrNotFound = errors.New("记录不存在")

func loadRecord() error {
    return fmt.Errorf("加载用户记录失败: %w", ErrNotFound)
}

func main() {
    err := loadRecord()
    if errors.Is(err, ErrNotFound) {
        fmt.Println("记录不存在")
    }
}
~~~

errors.Is 会沿着包装链判断是否包含 ErrNotFound。
