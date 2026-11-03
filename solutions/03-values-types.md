# 第 3 章参考答案：变量与类型

## 练习 3.1

~~~go
name := "小明"
age := 20
isStudent := true

fmt.Println(name, age, isStudent)
~~~

## 练习 3.2

~~~go
price := 100
taxRate := 0.08
total := float64(price) * (1 + taxRate)

fmt.Printf("含税价格：%.2f\n", total)
~~~

先把 price 从 int 转成 float64，再参与浮点数运算。

## 练习 3.3

~~~go
score := 60
score = 80
fmt.Println(score)
~~~

第一次声明使用 :=，之后修改已有变量使用 =。
