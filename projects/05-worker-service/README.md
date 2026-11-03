# 项目 5：并发 Worker 服务

## 目标

实现一个通用任务处理器：把任务放入队列，由固定数量的 worker 并发处理，最后收集结果。

本项目用于学习并发编程，不涉及自动化访问第三方账户或绕过服务限制。

## 任务模型

~~~go
type Job struct {
    ID    int
    Value int
}

type Result struct {
    JobID int
    Value int
    Err   error
}
~~~

## 功能要求

- worker 数量可配置。
- 任务通过 channel 分发。
- 结果通过 channel 返回。
- 支持 context 取消。
- 所有 goroutine 都能正常结束。
- 发生错误时能在结果中体现。
- 没有未保护的共享 map 写入。
- 有并发安全测试。

## 推荐目录

~~~text
worker-service/
├── main.go
├── worker.go
├── worker_test.go
└── README.md
~~~

## 推荐实现顺序

1. 先写一个同步的 process(Job) Result 函数。
2. 用一个 goroutine 执行一个任务。
3. 把任务输入改成 jobs channel。
4. 启动固定数量的 worker。
5. 用 WaitGroup 等待所有 worker 结束。
6. 关闭 jobs，等待 worker 退出。
7. 收集结果，并在所有 worker 完成后关闭 results。
8. 加入 context，让任务生产者和 worker 都能响应取消。
9. 用 go test -race ./... 检查数据竞争。
10. 测试 0 个任务、1 个任务、很多任务和取消场景。

## 生命周期图

~~~mermaid
flowchart TD
    A["创建 context 和 channels"] --> B["启动 N 个 worker"]
    B --> C["生产者提交任务"]
    C --> D["worker 从 jobs 读取"]
    D --> E["处理任务"]
    E --> F["写入 results"]
    F --> G{"还有任务？"}
    G -- "有" --> D
    G -- "没有" --> H["关闭 jobs"]
    H --> I["等待 worker 退出"]
    I --> J["关闭 results"]
    J --> K["收集并输出结果"]
~~~

## 并发安全提醒

- 不能在多个 goroutine 中无保护地同时写普通 map。
- WaitGroup 的 Add 应在对应 goroutine 启动前完成。
- 关闭 channel 的责任应该清晰，通常由发送方或任务协调者关闭。
- 不要重复关闭 channel。
- 不要让发送方在无人接收时永久阻塞。
- context 取消后，任务函数必须检查取消信号并及时退出。

## 验收清单

- [ ] worker 数量可配置。
- [ ] 所有任务都能得到结果。
- [ ] 没有任务时能正常结束。
- [ ] context 取消后能够退出。
- [ ] channel 关闭顺序正确。
- [ ] go test -race ./... 通过。
- [ ] 核心逻辑有测试。
