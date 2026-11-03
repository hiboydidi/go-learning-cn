# 第 4 章：条件判断与循环

## 学习目标

- 用 if / else 处理不同情况。
- 用 switch 处理多个分支。
- 用 for 重复执行代码。
- 理解 break 和 continue。

## 一、if：如果条件成立，就执行

~~~go
score := 82

if score >= 60 {
    fmt.Println("通过")
} else {
    fmt.Println("未通过")
}
~~~

Go 的 if 条件不需要括号，但花括号必须保留。

可以在 if 中先声明一个临时变量：

~~~go
if age := 20; age >= 18 {
    fmt.Println("成年人")
} else {
    fmt.Println("未成年人")
}
~~~

age 只在这个 if / else 结构中有效。

## 二、多个条件

~~~go
score := 82

if score >= 90 {
    fmt.Println("优秀")
} else if score >= 60 {
    fmt.Println("及格")
} else {
    fmt.Println("需要补习")
}
~~~

判断顺序从上到下。一旦某个条件成立，就执行对应分支，后面的分支不再执行。

## 三、switch：多个明确分支

~~~go
day := 2

switch day {
case 1:
    fmt.Println("星期一")
case 2:
    fmt.Println("星期二")
case 3:
    fmt.Println("星期三")
default:
    fmt.Println("其他日期")
}
~~~

Go 的 switch 默认不会自动贯穿到下一个 case。大多数时候不需要写 break。

## 四、for：重复执行

Go 只有 for 这一种循环关键字，但可以表达常见的循环形式。

### 形式 1：计数循环

~~~go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
~~~

执行过程：

1. i := 0：创建计数器。
2. i < 5：检查条件。
3. 执行循环体。
4. i++：计数器加 1。
5. 回到条件检查，直到条件不成立。

### 形式 2：条件循环（类似 while）

~~~go
count := 0

for count < 3 {
    fmt.Println("执行一次")
    count++
}
~~~

### 形式 3：无限循环

~~~go
for {
    fmt.Println("持续运行")
    break
}
~~~

如果没有 break、return 或其他退出条件，无限循环会一直运行。

## 五、break 和 continue

- break：结束当前循环。
- continue：跳过本轮剩余代码，直接进入下一轮循环。

~~~go
for i := 1; i <= 5; i++ {
    if i == 2 {
        continue
    }
    if i == 4 {
        break
    }
    fmt.Println(i)
}
~~~

输出：

~~~text
1
3
~~~

## 六、流程图理解分支

~~~mermaid
flowchart TD
    A["开始"] --> B["读取 score"]
    B --> C{"score >= 90 ?"}
    C -- "是" --> D["输出：优秀"]
    C -- "否" --> E{"score >= 60 ?"}
    E -- "是" --> F["输出：及格"]
    E -- "否" --> G["输出：未通过"]
    D --> H["结束"]
    F --> H
    G --> H
~~~

## 七、常见错误

1. 忘记更新循环变量，导致循环永远不结束。
2. 把 == 写成 =。前者比较，后者赋值。
3. 条件边界写错，例如想包含 5，却写成 i < 5。
4. 把 continue 理解成结束整个循环；结束整个循环的是 break。

## 本章练习

**练习 4.1**：定义 age。如果 age 大于等于 18，输出“可以进入”；否则输出“暂不允许”。

**练习 4.2**：定义 score，根据 90、60 两条边界输出优秀、及格或不及格。

**练习 4.3**：使用 for 输出 1 到 10。

**练习 4.4**：输出 1 到 10，但跳过 3，并在遇到 8 时结束循环。

<details>
<summary>参考答案：练习 4.4</summary>

~~~go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 3 {
            continue
        }
        if i == 8 {
            break
        }
        fmt.Println(i)
    }
}
~~~

</details>

**验收标准**：你能画出一个 if/else 流程图，并解释循环每次如何回到条件判断。
