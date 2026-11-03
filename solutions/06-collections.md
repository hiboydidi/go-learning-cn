# 第 6 章参考答案：数组、切片与映射

## 练习 6.1

~~~go
numbers := []int{1, 2, 3, 4, 5}
for _, number := range numbers {
    fmt.Println(number)
}
~~~

## 练习 6.2

~~~go
numbers := []int{2, 4, 6, 8, 10}
total := 0

for _, number := range numbers {
    total += number
}
fmt.Println(total)
~~~

结果是 30。

## 练习 6.3

~~~go
ages := map[string]int{
    "小明": 20,
    "小红": 22,
    "小刚": 19,
}

age, ok := ages["小红"]
if ok {
    fmt.Println(age)
} else {
    fmt.Println("没有找到")
}
~~~

## 练习 6.4

~~~go
words := []string{"go", "js", "go", "go", "python"}
counts := make(map[string]int)

for _, word := range words {
    counts[word]++
}

fmt.Println(counts)
~~~

结果：go 为 3，js 为 1，python 为 1。map 的打印顺序不保证固定。
