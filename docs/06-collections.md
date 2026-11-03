# 第 6 章：数组、切片与映射

## 学习目标

- 理解数组和切片的区别。
- 会追加、遍历和修改切片。
- 会使用 map 保存键值对。
- 能处理一组数据。

## 一、数组 array

数组是一组固定长度、相同类型的数据。

~~~go
var scores [3]int
scores[0] = 80
scores[1] = 90
scores[2] = 100

fmt.Println(scores)
~~~

数组长度属于类型的一部分。[3]int 和 [4]int 是不同的类型。

## 二、切片 slice

切片是 Go 中最常用的序列类型。它可以动态增长。

~~~go
scores := []int{80, 90, 100}
scores = append(scores, 95)

fmt.Println(scores)
fmt.Println(len(scores))
~~~

- len：当前元素数量。
- cap：底层数组从切片起点开始的容量。
- append：追加元素，必要时会分配新的底层数组。

也可以使用 make：

~~~go
names := make([]string, 0, 10)
names = append(names, "小明")
~~~

这表示长度为 0、容量为 10 的切片。

## 三、遍历切片

~~~go
scores := []int{80, 90, 100}

for index, score := range scores {
    fmt.Println(index, score)
}
~~~

如果不需要索引，可以使用下划线丢弃：

~~~go
for _, score := range scores {
    fmt.Println(score)
}
~~~

## 四、切片的注意点

切片不是一个独立复制出来的数组。不同切片可能共享底层数组。

~~~go
numbers := []int{1, 2, 3}
part := numbers[:2]
part[0] = 99

fmt.Println(numbers) // [99 2 3]
~~~

如果需要独立复制，可以使用：

~~~go
copied := append([]int(nil), numbers...)
~~~

## 五、映射 map

map 保存键值对，适合通过 key 快速查找 value。

~~~go
ages := map[string]int{
    "小明": 20,
    "小红": 22,
}

ages["小刚"] = 19
fmt.Println(ages["小明"])
~~~

判断某个 key 是否存在：

~~~go
age, ok := ages["小李"]
if ok {
    fmt.Println("年龄：", age)
} else {
    fmt.Println("没有找到")
}
~~~

删除：

~~~go
delete(ages, "小红")
~~~

注意：map 的遍历顺序不保证固定。不要依赖它每次都按相同顺序输出。

## 六、用 map 统计词频

~~~go
words := []string{"go", "python", "go", "java", "go"}
counts := make(map[string]int)

for _, word := range words {
    counts[word]++
}

fmt.Println(counts["go"]) // 3
~~~

## 七、何时使用哪一种

| 类型 | 适合场景 |
|---|---|
| array | 长度固定、类型明确 |
| slice | 列表、批量数据、动态集合 |
| map | 通过 key 查找和统计 |

## 本章练习

**练习 6.1**：创建一个包含 5 个整数的切片，遍历并打印每个元素。

**练习 6.2**：计算切片中所有数字的总和。

**练习 6.3**：创建一个 map，保存 3 个人的姓名和年龄，然后查找其中一个人。

**练习 6.4**：统计 []string{"go", "js", "go", "go", "python"} 中每个词出现的次数。

<details>
<summary>参考答案：练习 6.2</summary>

~~~go
numbers := []int{2, 4, 6, 8, 10}
total := 0

for _, number := range numbers {
    total += number
}

fmt.Println(total)
~~~

</details>

**验收标准**：能解释数组长度固定、切片可增长，以及 map 的 key/value 结构。
