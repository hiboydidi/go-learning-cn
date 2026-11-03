# 第 3 章：变量、常量与基本类型

## 学习目标

- 使用变量保存数据。
- 认识 Go 的基本类型。
- 理解类型转换和零值。
- 看懂 := 和 var 的区别。

## 变量是什么

变量可以先理解为一个有名字的储物格。你给它一个名字和类型，然后把值放进去。

~~~go
var username string = "小明"
var age int = 20

fmt.Println(username)
fmt.Println(age)
~~~

## 声明变量的三种常见方式

~~~go
var name string = "Go"
var city = "Jakarta"
count := 3
~~~

- var name string = "Go"：显式写出类型和值。
- var city = "Jakarta"：让 Go 根据初始值推断类型。
- count := 3：短变量声明，只能在函数内部使用。

修改已有变量时使用 =，不要再次使用 :=：

~~~go
count := 3
count = 4
~~~

## 基本类型

| 类型 | 用途 | 示例 |
|---|---|---|
| string | 文本 | "hello" |
| bool | 真或假 | true |
| int | 整数 | 42 |
| int64 | 64 位整数 | 9000 |
| float64 | 浮点数 | 3.14 |
| byte | uint8 的别名，常用于字节 | 'A' |
| rune | int32 的别名，常用于 Unicode 字符 | '中' |

## 零值

变量声明后没有显式赋值，Go 会给它零值。

| 类型 | 零值 |
|---|---|
| int、float | 0 |
| bool | false |
| string | 空字符串 "" |
| 指针、slice、map、channel、函数、接口 | nil |

示例：

~~~go
var count int
var active bool
var title string

fmt.Println(count)  // 0
fmt.Println(active) // false
fmt.Println(title)  // 空字符串
~~~

## 常量

常量使用 const 声明，声明后不能重新赋值。

~~~go
const appName = "学习计时器"
const maxRetries = 3
~~~

## 类型转换

Go 不会自动把所有数字类型互相转换，需要显式转换：

~~~go
var a int = 10
var b float64 = float64(a)

fmt.Println(b)
~~~

注意：把浮点数转换成整数会丢掉小数部分。

## 输入与输出

fmt.Println 适合输出。初学阶段可用 fmt.Scan 读取简单输入：

~~~go
var name string
fmt.Print("请输入名字：")
fmt.Scan(&name)
fmt.Println("你好，", name)
~~~

&name 表示把变量的地址交给 Scan，让它能把输入写入变量。指针会在第 7 章进一步解释。

## 常见错误

1. 把 := 当成普通赋值符号使用。
2. 给 int 变量赋字符串。
3. 声明变量后完全没有使用，导致编译错误。
4. 以为变量未赋值时是随机值；Go 会提供零值。

## 本章练习

**练习 3.1**：声明 name、age、isStudent 三个变量，并打印它们。

**练习 3.2**：声明一个整数 price 和一个浮点数 taxRate，计算含税价格。

**练习 3.3**：声明一个变量 score，先赋值 60，再改成 80，最后打印。

<details>
<summary>参考答案：练习 3.2</summary>

~~~go
package main

import "fmt"

func main() {
    price := 100
    taxRate := 0.08
    total := float64(price) * (1 + taxRate)
    fmt.Printf("含税价格：%.2f\n", total)
}
~~~

</details>

**验收标准**：你能解释 var、:=、=、const 和零值的区别。
