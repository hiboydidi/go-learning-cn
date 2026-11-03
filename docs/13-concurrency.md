# 第 13 章：并发——goroutine、channel 与 context

## 学习目标

- 理解并发和并行的基本区别。
- 使用 goroutine 启动并发任务。
- 使用 channel 传递结果。
- 使用 WaitGroup 等待任务结束。
- 使用 context 取消任务。
- 避免常见数据竞争。

## 一、并发是什么

并发表示多个任务在一段时间内交错推进；并行表示多个任务在同一时刻真正同时执行。

Go 使用 goroutine 表达轻量级并发任务。

~~~go
go doWork()
~~~

go 关键字会启动一个 goroutine。它不会自动等待函数结束。

## 二、为什么只写 go 不够

~~~go
go fmt.Println("后台任务")
fmt.Println("主程序结束")
~~~

主函数结束时，程序也可能结束，后台任务还没来得及完成。

需要一种方式等待任务完成。

## 三、WaitGroup

~~~go
var wg sync.WaitGroup

wg.Add(2)

go func() {
    defer wg.Done()
    fmt.Println("任务 A")
}()

go func() {
    defer wg.Done()
    fmt.Println("任务 B")
}()

wg.Wait()
fmt.Println("全部完成")
~~~

- Add：增加等待计数。
- Done：完成一个任务，计数减一。
- Wait：等待计数变成 0。

## 四、channel：在 goroutine 之间传递数据

~~~go
results := make(chan int)

go func() {
    results <- 42
}()

value := <-results
fmt.Println(value)
~~~

- ch <- value：发送数据。
- value := <-ch：接收数据。
- make(chan int)：创建无缓冲 channel。

无缓冲 channel 的发送和接收需要彼此配合。缓冲 channel 可以容纳指定数量的元素：

~~~go
jobs := make(chan int, 5)
jobs <- 1
jobs <- 2
~~~

## 五、select

select 可以等待多个 channel 操作中的一个：

~~~go
select {
case result := <-results:
    fmt.Println("结果：", result)
case <-time.After(2 * time.Second):
    fmt.Println("等待超时")
}
~~~

## 六、context 取消任务

~~~go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("任务收到取消信号")
            return
        default:
            doSmallWork()
        }
    }
}()

cancel()
~~~

真实代码不要在 default 分支里无休止地空转。应当执行实际工作、等待输入，或使用定时器控制频率。

## 七、互斥锁 Mutex

多个 goroutine 同时修改同一个变量时，可能产生数据竞争。

~~~go
var (
    mu      sync.Mutex
    counter int
)

mu.Lock()
counter++
mu.Unlock()
~~~

锁必须在所有访问共享状态的路径上正确使用。更简单的设计通常是让 goroutine 通过 channel 传递数据，减少共享可变状态。

## 八、Worker Pool 工作池

工作池通常由三部分组成：

1. jobs channel：等待执行的任务。
2. 固定数量的 worker：消费任务。
3. results channel：收集执行结果。

~~~mermaid
flowchart LR
    A["任务生产者"] --> B["jobs channel"]
    B --> C["Worker 1"]
    B --> D["Worker 2"]
    B --> E["Worker 3"]
    C --> F["results channel"]
    D --> F
    E --> F
    F --> G["结果收集器"]
~~~

完整示例见 examples/worker-pool。

## 九、并发安全检查

~~~bash
go test -race ./...
~~~

race detector 可以帮助发现运行过程中触发的数据竞争。它不能证明程序绝对没有竞争，但很适合在测试阶段使用。

## 本章练习

**练习 13.1**：启动两个 goroutine，分别打印不同任务名称，并等待它们结束。

**练习 13.2**：创建一个 channel，让 goroutine 发送 100，主 goroutine 接收并打印。

**练习 13.3**：创建 3 个 worker，处理 10 个整数任务。

**练习 13.4**：为一个长时间运行的任务增加 context 取消能力。

**验收标准**：能解释 goroutine、channel、WaitGroup、Mutex 和 context 各自解决什么问题。
