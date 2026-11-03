# 第 9 章参考答案：接口与泛型

## 练习 9.1 和 9.2

~~~go
type Runner interface {
    Run() string
}

type Car struct{}

func (Car) Run() string {
    return "汽车正在行驶"
}

type Robot struct{}

func (Robot) Run() string {
    return "机器人正在执行任务"
}

func start(r Runner) {
    fmt.Println(r.Run())
}
~~~

调用 start(Car{}) 和 start(Robot{})。两个类型都实现了 Run 方法，因此都满足 Runner。

## 练习 9.3

~~~go
func Last[T any](items []T) (T, bool) {
    var zero T

    if len(items) == 0 {
        return zero, false
    }

    return items[len(items)-1], true
}
~~~

空切片没有最后一个元素，因此返回 T 的零值和 false。
