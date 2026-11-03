# 第 11 章：文件与 JSON

## 学习目标

- 读取和写入本地文件。
- 使用 bufio 逐行读取文本。
- 使用 encoding/json 编码和解码 JSON。
- 知道如何处理文件操作错误。

## 一、读取整个文件

~~~go
data, err := os.ReadFile("notes.txt")
if err != nil {
    return err
}

fmt.Println(string(data))
~~~

ReadFile 会把文件内容读入内存，适合文件不大的场景。

## 二、写入文件

~~~go
content := []byte("今天学习 Go\n")
err := os.WriteFile("notes.txt", content, 0644)
if err != nil {
    return err
}
~~~

0644 是 Unix 文件权限表示。不要把不可信输入当成任意文件路径。

## 三、逐行读取

大文件不适合一次性全部读入内存，可以使用 bufio.Scanner：

~~~go
file, err := os.Open("app.log")
if err != nil {
    return err
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
}

if err := scanner.Err(); err != nil {
    return err
}
~~~

defer 会在当前函数返回前执行，适合关闭文件等清理动作。

## 四、JSON 是什么

JSON 是一种常见的数据交换格式：

~~~json
{
  "name": "小明",
  "age": 20,
  "active": true
}
~~~

Go 使用 encoding/json 标准库。

## 五、把结构体编码成 JSON

~~~go
type User struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Active bool   `json:"active"`
}

user := User{Name: "小明", Age: 20, Active: true}

data, err := json.MarshalIndent(user, "", "  ")
if err != nil {
    return err
}

fmt.Println(string(data))
~~~

结构体字段通常需要导出（首字母大写），encoding/json 才能处理它们。

## 六、把 JSON 解码成结构体

~~~go
input := []byte("{\"name\":\"小红\",\"age\":22,\"active\":true}")

var user User
if err := json.Unmarshal(input, &user); err != nil {
    return err
}

fmt.Println(user.Name)
~~~

Unmarshal 需要一个指向目标变量的指针，因为它要把解析结果写入变量。

## 七、文件 + JSON：保存一个配置

常见流程：

1. 定义结构体。
2. 将结构体 Marshal 成 JSON。
3. 用 os.WriteFile 保存。
4. 用 os.ReadFile 读取。
5. 用 json.Unmarshal 还原结构体。
6. 每一步都检查错误。

## 本章练习

**练习 11.1**：创建一个 notes.txt 文件，写入三行文本，再读取并打印。

**练习 11.2**：定义 Task 结构体，字段包含 ID、Title、Done。

**练习 11.3**：把一个 Task 编码成 JSON 字符串，再解码回 Task。

**练习 11.4**：把多个 Task 保存到 tasks.json，再从文件读回来。

**验收标准**：能完成“结构体 → JSON → 文件 → 文件读取 → 结构体”的完整往返。
