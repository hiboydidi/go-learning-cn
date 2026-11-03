# 第 11 章参考答案：文件与 JSON

## 练习 11.1

~~~go
content := []byte("第一行\n第二行\n第三行\n")

if err := os.WriteFile("notes.txt", content, 0644); err != nil {
    return err
}

data, err := os.ReadFile("notes.txt")
if err != nil {
    return err
}

fmt.Println(string(data))
~~~

## 练习 11.2 和 11.3

~~~go
type Task struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}

task := Task{ID: 1, Title: "学习 JSON", Done: false}

data, err := json.Marshal(task)
if err != nil {
    return err
}

var decoded Task
if err := json.Unmarshal(data, &decoded); err != nil {
    return err
}
~~~

## 练习 11.4

~~~go
tasks := []Task{
    {ID: 1, Title: "学习 Go", Done: false},
    {ID: 2, Title: "写一个程序", Done: true},
}

data, err := json.MarshalIndent(tasks, "", "  ")
if err != nil {
    return err
}

if err := os.WriteFile("tasks.json", data, 0644); err != nil {
    return err
}

saved, err := os.ReadFile("tasks.json")
if err != nil {
    return err
}

var loaded []Task
if err := json.Unmarshal(saved, &loaded); err != nil {
    return err
}
~~~

完整项目还应处理文件不存在、文件损坏和写入失败等情况。
