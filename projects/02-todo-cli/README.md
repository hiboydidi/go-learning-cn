# 项目 2：Todo 命令行管理器

## 目标

编写一个命令行任务管理器，支持新增、查看、完成和删除任务，并将数据保存到本地 JSON 文件。

## 数据结构

可以从下面的结构体开始：

~~~go
type Task struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}
~~~

## 命令设计

~~~bash
todo add "学习 Go 切片"
todo list
todo done 1
todo delete 1
~~~

第一版可以使用简单的位置参数解析；熟悉后再研究 flag 或 cobra 等 CLI 工具。

## 推荐目录

~~~text
todo/
├── main.go
├── task.go
├── storage.go
└── tasks.json
~~~

## 推荐实现顺序

1. 定义 Task 结构体。
2. 用 []Task 在内存中保存任务。
3. 实现新增、列表、完成、删除四个操作。
4. 用 encoding/json 编码和解码任务列表。
5. 用 os.ReadFile 和 os.WriteFile 持久化数据。
6. 文件不存在时，按“空任务列表”处理。
7. 文件内容损坏时，返回明确错误，不要静默覆盖。
8. 给任务操作和存储逻辑编写测试。

## 验收清单

- [ ] 新增任务后能在列表中看到。
- [ ] 完成任务后 Done 为 true。
- [ ] 删除任务后列表中不再出现。
- [ ] 退出程序再启动，任务仍然存在。
- [ ] 不存在的任务 ID 会提示错误。
- [ ] JSON 文件损坏时能提示错误。
- [ ] 核心逻辑有测试。

## 进阶

- 支持任务优先级和截止日期。
- 支持按完成状态筛选。
- 支持导出 CSV。
- 将存储逻辑抽象成接口，但只有在确实需要多个存储实现时再做。
