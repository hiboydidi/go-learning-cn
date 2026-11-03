# 第 17 章参考答案：数据库

## 练习 17.1

sql.Open 创建并返回数据库句柄，通常代表一个连接池；它不保证立即建立真实连接。需要用 Ping 或 PingContext 检查数据库是否可连接。

## 练习 17.2

~~~go
var name string

err := db.QueryRowContext(
    ctx,
    "SELECT name FROM users WHERE id = $1",
    userID,
).Scan(&name)

if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
        return fmt.Errorf("用户不存在")
    }
    return err
}
~~~

占位符 $1 适用于常见 PostgreSQL 驱动，其他驱动可能使用不同占位符。

## 练习 17.3

直接拼接输入会让输入内容改变 SQL 语句结构，可能造成 SQL 注入。参数化查询把 SQL 结构和参数值分开交给驱动处理。

## 练习 17.4

事务步骤：

1. BeginTx 开启事务。
2. 执行扣款 SQL。
3. 检查扣款是否成功以及余额是否足够。
4. 执行入账 SQL。
5. 任一步失败就 Rollback。
6. 全部成功后 Commit。

扣款和入账必须在同一个事务中，不能一边成功另一边失败。
