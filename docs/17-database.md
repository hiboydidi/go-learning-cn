# 第 17 章：数据库与 database/sql

## 学习目标

- 理解 Go 如何连接关系型数据库。
- 认识 database/sql 的职责。
- 会写查询、插入和事务的基本代码。
- 知道连接池和 SQL 参数化的重要性。

## 一、database/sql 是什么

database/sql 是 Go 标准库提供的数据库通用接口。它定义了连接池、查询、执行语句和事务等能力，但它本身不是数据库驱动。

使用具体数据库时，还需要安装对应驱动，例如 PostgreSQL、MySQL 或 SQLite 驱动。驱动的安装方式、连接字符串和占位符格式，以对应驱动的官方文档为准。

## 二、连接数据库

概念示例：

~~~go
import (
    "context"
    "database/sql"
    "time"
)

// 需要根据你选择的数据库安装并导入对应 driver。
db, err := sql.Open("driver-name", "your-connection-string")
if err != nil {
    return err
}
defer db.Close()

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := db.PingContext(ctx); err != nil {
    return err
}
~~~

sql.Open 通常先创建连接池句柄，不一定立即建立真实连接。PingContext 用于检查数据库是否可连接。

不要把真实数据库密码提交到 Git。连接信息应从环境变量或安全的密钥管理系统读取。

## 三、连接池

database/sql 管理连接池。常见配置：

~~~go
db.SetMaxOpenConns(20)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(30 * time.Minute)
~~~

这些数字只是示例，不是所有项目都适用的固定最佳值。需要根据数据库容量、并发量和查询耗时调整。

## 四、查询一条记录

~~~go
var name string
var age int

err := db.QueryRowContext(
    ctx,
    "SELECT name, age FROM users WHERE id = $1",
    userID,
).Scan(&name, &age)

if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
        return fmt.Errorf("用户不存在")
    }
    return err
}
~~~

注意：SQL 占位符因驱动而异。PostgreSQL 常见 $1，部分 SQLite/MySQL 驱动使用 ?。应查看具体驱动文档。

## 五、查询多条记录

~~~go
rows, err := db.QueryContext(ctx, "SELECT id, name FROM users")
if err != nil {
    return err
}
defer rows.Close()

for rows.Next() {
    var id int
    var name string

    if err := rows.Scan(&id, &name); err != nil {
        return err
    }

    fmt.Println(id, name)
}

if err := rows.Err(); err != nil {
    return err
}
~~~

rows 使用后要关闭，并检查 rows.Err()。

## 六、插入或更新

~~~go
result, err := db.ExecContext(
    ctx,
    "INSERT INTO users(name, age) VALUES($1, $2)",
    name,
    age,
)
if err != nil {
    return err
}

id, err := result.LastInsertId()
~~~

并不是所有数据库驱动都支持 LastInsertId。PostgreSQL 常用 RETURNING 子句读取新 ID。

## 七、事务 transaction

当多个数据库操作必须一起成功或一起失败时，使用事务。

~~~go
tx, err := db.BeginTx(ctx, nil)
if err != nil {
    return err
}
defer tx.Rollback()

if _, err := tx.ExecContext(ctx, queryA, argsA...); err != nil {
    return err
}

if _, err := tx.ExecContext(ctx, queryB, argsB...); err != nil {
    return err
}

if err := tx.Commit(); err != nil {
    return err
}
~~~

如果函数提前返回，defer Rollback 会清理未提交事务；事务提交后再 Rollback 不会撤销已提交内容。

## 八、安全注意事项

不要这样拼接 SQL：

~~~go
query := "SELECT * FROM users WHERE name = '" + userInput + "'"
~~~

应该使用参数化查询：

~~~go
row := db.QueryRowContext(ctx, "SELECT * FROM users WHERE name = $1", userInput)
~~~

参数化查询可以避免把用户输入直接拼进 SQL 语句。表名、列名等不能简单当作普通参数传入，动态标识符需要白名单控制。

## 本章练习

**练习 17.1**：解释 sql.Open 为什么不一定立即建立真实连接。

**练习 17.2**：写出查询一条用户记录的伪代码，并处理 sql.ErrNoRows。

**练习 17.3**：说明为什么不应该把用户输入直接拼接到 SQL 字符串中。

**练习 17.4**：设计一个转账事务：扣款和入账必须一起成功或一起失败。

**验收标准**：能解释 database/sql、驱动、连接池、参数化查询和事务的职责。
