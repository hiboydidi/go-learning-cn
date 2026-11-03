# 第 5 章：函数

## 学习目标

- 理解函数为什么存在。
- 会定义参数和返回值。
- 能拆分一个较大的任务。
- 看懂多返回值和可变参数。

## 函数解决什么问题

假设程序中有 10 个地方都需要计算含税价格。如果每个地方都复制一遍公式，之后修改税率就要改 10 次。

把计算过程写成函数，就可以重复调用：

~~~go
func totalPrice(price float64, taxRate float64) float64 {
    return price * (1 + taxRate)
}
~~~

调用：

~~~go
total := totalPrice(100, 0.08)
fmt.Println(total)
~~~

## 函数的结构

~~~go
func 函数名(参数列表) 返回值类型 {
    // 函数体
    return 返回值
}
~~~

参数是函数接收的数据，返回值是函数计算完成后交还的数据。

## 多个参数和多个返回值

~~~go
func divide(a float64, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("除数不能为 0")
    }
    return a / b, nil
}
~~~

调用：

~~~go
result, err := divide(10, 2)
if err != nil {
    fmt.Println("计算失败：", err)
    return
}
fmt.Println(result)
~~~

Go 常用多返回值表达“结果 + 错误”。错误处理在第 10 章详细学习。

## 参数类型可以合并书写

~~~go
func add(a, b int) int {
    return a + b
}
~~~

## 没有返回值的函数

~~~go
func printWelcome(name string) {
    fmt.Println("欢迎你，", name)
}
~~~

## 可变参数

当参数数量不确定时，可以使用 ...：

~~~go
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}
~~~

调用：

~~~go
fmt.Println(sum(1, 2))
fmt.Println(sum(1, 2, 3, 4))
~~~

## 函数拆解流程

~~~mermaid
flowchart TD
    A["主程序接收输入"] --> B["调用 validate"]
    B --> C{"输入有效？"}
    C -- "否" --> D["返回错误"]
    C -- "是" --> E["调用 calculate"]
    E --> F["返回计算结果"]
    F --> G["主程序打印结果"]
~~~

## 变量作用域

- 函数内部声明的局部变量通常只能在函数内部使用。
- 函数参数只在函数调用期间可用。
- 包级变量在同一个 package 的文件中可见。
- 首字母大写的包级名称可以被其他 package 引用。

初学阶段尽量把变量放在最小需要的范围内。

## 本章练习

**练习 5.1**：写一个 add 函数，接收两个整数，返回它们的和。

**练习 5.2**：写一个 isEven 函数，接收整数，返回 bool，判断是否为偶数。

**练习 5.3**：写一个 max 函数，接收两个整数，返回较大的那个。

**练习 5.4**：写一个 divide 函数，除数为 0 时返回错误。

<details>
<summary>参考答案：练习 5.1 和 5.2</summary>

~~~go
func add(a, b int) int {
    return a + b
}

func isEven(number int) bool {
    return number%2 == 0
}
~~~

</details>

**验收标准**：能把“读取输入、验证输入、计算、输出”拆成不同函数。
