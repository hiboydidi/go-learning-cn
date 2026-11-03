# 第 13 章参考答案：并发

## 练习 13.1

~~~go
var wg sync.WaitGroup

for _, name := range []string{"任务 A", "任务 B"} {
    wg.Add(1)

    go func(taskName string) {
        defer wg.Done()
        fmt.Println(taskName)
    }(name)
}

wg.Wait()
~~~

把循环变量作为参数传入 goroutine，能明确每个任务使用的值。

## 练习 13.2

~~~go
result := make(chan int)

go func() {
    result <- 100
}()

value := <-result
fmt.Println(value)
~~~

## 练习 13.3

可以直接运行 examples/worker-pool：

~~~bash
go run ./examples/worker-pool
~~~

其中 workerCount 控制 worker 数量，jobCount 控制任务数量。

## 练习 13.4

~~~go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

done := make(chan struct{})

go func() {
    defer close(done)

    for {
        select {
        case <-ctx.Done():
            return
        default:
            // 执行一小段可中断的工作
            time.Sleep(100 * time.Millisecond)
        }
    }
}()

cancel()
<-done
~~~

真实工作函数也应在合适的位置检查 ctx.Done()，不要只在循环入口检查。
