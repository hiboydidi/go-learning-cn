# 第 4 章参考答案：条件与循环

## 练习 4.1

~~~go
age := 20
if age >= 18 {
    fmt.Println("可以进入")
} else {
    fmt.Println("暂不允许")
}
~~~

## 练习 4.2

~~~go
score := 82

if score >= 90 {
    fmt.Println("优秀")
} else if score >= 60 {
    fmt.Println("及格")
} else {
    fmt.Println("不及格")
}
~~~

## 练习 4.3

~~~go
for i := 1; i <= 10; i++ {
    fmt.Println(i)
}
~~~

## 练习 4.4

~~~go
for i := 1; i <= 10; i++ {
    if i == 3 {
        continue
    }
    if i == 8 {
        break
    }
    fmt.Println(i)
}
~~~

输出 1、2、4、5、6、7。continue 跳过 3，break 在 8 时结束整个循环。
