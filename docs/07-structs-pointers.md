# 第 7 章：字符串、结构体、指针与方法

## 学习目标

- 理解 byte、rune 和字符串。
- 用 struct 表示一个业务对象。
- 理解指针的基本用途。
- 给结构体定义方法。

## 一、字符串、byte 和 rune

Go 的 string 是一段不可变的字节序列。UTF-8 中文字符通常占多个字节。

~~~go
text := "Go语言"

fmt.Println(len(text)) // 字节数，不是字符数

for index, r := range text {
    fmt.Println(index, string(r))
}
~~~

- byte：一个字节，适合处理原始字节。
- rune：一个 Unicode 码点，适合按字符遍历。
- len(string)：返回字节数。
- utf8.RuneCountInString：统计 rune 数量。

~~~go
import (
    "fmt"
    "unicode/utf8"
)

text := "你好"
fmt.Println(len(text))
fmt.Println(utf8.RuneCountInString(text))
~~~

常用字符串工具在 strings 包中，例如 Contains、Split、TrimSpace、ToLower。

## 二、结构体 struct

结构体把多个相关字段组合成一个类型。

~~~go
type User struct {
    Name  string
    Age   int
    Email string
}
~~~

创建：

~~~go
user := User{
    Name:  "小明",
    Age:   20,
    Email: "ming@example.com",
}

fmt.Println(user.Name)
~~~

可以把 User 理解为一张数据表的单行记录：一个对象由多个字段组成。

## 三、指针 pointer

指针保存的是另一个变量的内存地址。

~~~go
number := 10
pointer := &number

fmt.Println(*pointer) // 读取指针指向的值

*pointer = 20
fmt.Println(number)   // 20
~~~

- &number：取得变量地址。
- *pointer：访问指针指向的值。
- *int：表示“指向 int 的指针类型”。

指针为 nil 时没有指向有效对象，解引用 nil 会导致 panic。

## 四、为什么函数有时需要指针参数

如果函数只接收普通值，通常操作的是一份值拷贝：

~~~go
func changeValue(number int) {
    number = 100
}
~~~

调用后，外部变量不会被修改。

如果需要修改调用方的变量，可以传指针：

~~~go
func changeValue(number *int) {
    *number = 100
}
~~~

## 五、方法 method

方法是绑定到某个类型上的函数。

~~~go
type Account struct {
    Balance float64
}

func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}
~~~

调用：

~~~go
account := Account{Balance: 100}
account.Deposit(50)
fmt.Println(account.Balance) // 150
~~~

这里使用指针接收者，因为 Deposit 需要修改原对象。

## 六、值接收者还是指针接收者

- 值接收者：方法操作的是接收者的值拷贝，适合不需要修改对象的场景。
- 指针接收者：方法可以修改原对象，也避免复制较大的结构体。

同一个类型的方法通常保持一致的接收者风格，除非有明确原因。

## 本章练习

**练习 7.1**：定义 Product 结构体，包含 Name、Price、Stock 三个字段。

**练习 7.2**：创建一个 Product 实例，打印商品名称和价格。

**练习 7.3**：编写一个函数，用指针把库存增加 5。

**练习 7.4**：给 Product 添加一个方法，计算总库存价值（Price × Stock）。

<details>
<summary>参考答案：练习 7.3</summary>

~~~go
type Product struct {
    Name  string
    Price float64
    Stock int
}

func addStock(product *Product, amount int) {
    product.Stock += amount
}
~~~

</details>

**验收标准**：能解释 struct 用来组织什么数据，以及 & 和 * 的基本含义。
