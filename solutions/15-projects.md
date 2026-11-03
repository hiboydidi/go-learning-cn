# 第 15 章参考答案：实战项目设计

项目允许有不同实现。下面是建议的模块边界和验收思路，不是唯一正确答案。

## 项目 1：计算器

建议函数：

- readInput：读取数字和运算符。
- calculate：根据运算符计算并返回 error。
- run：控制循环和退出。

核心逻辑应能脱离终端输入单独测试。

## 项目 2：Todo CLI

建议结构：

- Task：任务数据结构。
- Store：负责 JSON 文件读取和写入。
- Add、List、Complete、Delete：任务业务逻辑。
- main：解析命令行参数并调用业务逻辑。

先实现内存版本，再增加文件持久化。

## 项目 3：日志分析器

建议结构：

- AnalyzeLine：解析单行日志。
- Analyze：逐行统计。
- Report：组织输出结果。
- main：读取参数和文件。

把统计逻辑与文件读取分开，方便测试。

## 项目 4：API 客户端

建议结构：

- Client：封装 http.Client。
- GetJSON：发送请求并解析 JSON。
- APIError：保存 HTTP 状态码和错误信息。
- main：读取 URL、调用客户端、打印结果。

使用 httptest 测试，不依赖线上 API。

## 项目 5：Worker 服务

建议结构：

- Job：任务输入。
- Result：任务结果。
- Worker：处理任务。
- Pool：启动 worker、分发任务、收集结果。
- context：控制取消。
- tests：覆盖任务完成和取消。

所有项目共同验收要求：正常输入正确、错误输入可解释、核心逻辑有测试、README 能让别人运行。
