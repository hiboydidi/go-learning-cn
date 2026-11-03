# 第 10 章：错误处理

## 学习目标

- 理解 Go 为什么经常显式返回 error。
- 能创建、返回和检查错误。
- 理解错误包装。
- 区分 error、panic 和 recover。

## 一、error 是什么

程序运行时可能失败，例如：

- 文件不存在。
- 网络请求超时。
- 用户输入不符合要求。
- 除数为 0。
- 数据格式错误。

Go 通常把错误作为函数的一个返回值交给调用方处理。

~~~go
func findUser(id int) (string, error) {
    if id <= 0 {
        return "", fmt.Errorf("用户 ID 必须大于 0")
    }
    return "小明", nil
}
~~~

nil 表示没有错误。

## 二、检查错误

~~~go
name, err := findUser(0)
if err != nil {
    fmt.Println("查询失败：", err)
    return
}

fmt.Println(name)
~~~

习惯上先处理错误，再使用正常结果。

## 三、创建错误

~~~go
var ErrNotFound = errors.New("记录不存在")
~~~

当调用方需要判断某一类特定错误时，可以定义可比较的 sentinel error。

## 四、包装错误

~~~go
func loadConfig() error {
    _, err := os.ReadFile("config.json")
    if err != nil {
        return fmt.Errorf("读取配置失败: %w", err)
    }
    return nil
}
~~~

%w 会包装原始错误，保留错误链。

## 五、errors.Is 和 errors.As

errors.Is 用于判断错误链中是否包含某个特定错误：

~~~go
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("文件不存在")
}
~~~

errors.As 用于把错误链中的某种错误类型取出来：

~~~go
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("路径操作失败：", pathErr.Path)
}
~~~

## 六、panic 和 recover

- error：预期内、可以交给调用方处理的失败。
- panic：程序无法正常继续时触发的异常控制流。
- recover：只能在 defer 调用的函数中捕获当前 goroutine 的 panic。

普通业务失败优先返回 error，不要用 panic 代替正常错误处理。

~~~go
func safeRun() {
    defer func() {
        if value := recover(); value != nil {
            fmt.Println("捕获到 panic：", value)
        }
    }()

    panic("演示用 panic")
}
~~~

recover 不是通用错误处理方案，也不能捕获其他 goroutine 中的 panic。

## 七、错误处理流程

~~~mermaid
flowchart TD
    A["调用函数"] --> B["接收 result 和 err"]
    B --> C{"err != nil ?"}
    C -- "是" --> D["记录上下文并处理或返回"]
    C -- "否" --> E["使用 result"]
    D --> F["结束当前流程"]
    E --> F
~~~

## 本章练习

**练习 10.1**：写一个除法函数，除数为 0 时返回 error。

**练习 10.2**：定义 ErrInvalidAge，当年龄小于 0 时返回它。

**练习 10.3**：用 fmt.Errorf 的 %w 包装一个错误，再用 errors.Is 判断原始错误。

**验收标准**：能说明为什么 Go 代码经常先检查 err，再处理结果。
