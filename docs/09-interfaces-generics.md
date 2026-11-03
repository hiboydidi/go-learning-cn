# 第 9 章：接口与泛型

## 学习目标

- 理解 interface 描述的是“能力”。
- 知道 Go 接口采用隐式实现。
- 能使用简单泛型函数。
- 不为了抽象而抽象。

## 一、接口是什么

接口定义一组方法。一个类型只要实现了接口要求的方法，就满足这个接口，不需要显式声明 implements。

~~~go
type Speaker interface {
    Speak() string
}

type Dog struct{}

func (Dog) Speak() string {
    return "汪"
}

type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "你好，我是" + p.Name
}
~~~

Dog 和 Person 都满足 Speaker 接口。

~~~go
func introduce(s Speaker) {
    fmt.Println(s.Speak())
}
~~~

调用：

~~~go
introduce(Dog{})
introduce(Person{Name: "小明"})
~~~

introduce 不关心传入的是 Dog 还是 Person，只关心它能不能 Speak。

## 二、接口的价值

接口让调用方依赖“需要的能力”，而不是依赖某个具体实现。

例如，业务只需要写入日志：

~~~go
type Logger interface {
    Println(values ...any)
}
~~~

真实项目中，接口应该从实际需求出发，尽量小而明确。

## 三、any 和空接口

any 是 interface{} 的别名，表示可以接收任意类型的值。

~~~go
func printValue(value any) {
    fmt.Printf("%T: %v\n", value, value)
}
~~~

但“能接收任何类型”不代表“应该到处使用 any”。能明确类型时，优先使用明确类型。

## 四、类型断言和类型 switch

~~~go
var value any = "hello"

text, ok := value.(string)
if ok {
    fmt.Println("这是字符串：", text)
}
~~~

也可以用类型 switch：

~~~go
switch v := value.(type) {
case string:
    fmt.Println("字符串：", v)
case int:
    fmt.Println("整数：", v)
default:
    fmt.Printf("其他类型：%T\n", v)
}
~~~

## 五、泛型是什么

泛型允许函数或类型使用类型参数，让同一段逻辑适用于多种类型，同时保留编译期类型检查。

~~~go
func First[T any](items []T) (T, bool) {
    var zero T
    if len(items) == 0 {
        return zero, false
    }
    return items[0], true
}
~~~

调用：

~~~go
name, ok := First([]string{"Go", "Python"})
number, found := First([]int{10, 20})
~~~

Go 会根据参数推断 T 的具体类型。

## 六、什么时候使用泛型

适合：

- 一段算法对多种类型都成立。
- 不希望为了不同类型复制多份相同逻辑。
- 希望复用代码但保留类型安全。

不适合：

- 只是为了让代码看起来高级。
- 不同类型实际上需要不同业务逻辑。
- 一个简单接口已经足够表达需求。

## 本章练习

**练习 9.1**：定义一个接口 Runner，要求有 Run() string 方法。

**练习 9.2**：定义 Car 和 Robot 两个类型，让它们都实现 Runner。

**练习 9.3**：写一个泛型函数 Last，返回切片最后一个元素和是否成功。

<details>
<summary>参考答案：练习 9.3</summary>

~~~go
func Last[T any](items []T) (T, bool) {
    var zero T
    if len(items) == 0 {
        return zero, false
    }
    return items[len(items)-1], true
}
~~~

</details>

**验收标准**：能解释接口关注方法能力，泛型关注可复用的类型参数。
