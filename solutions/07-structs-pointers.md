# 第 7 章参考答案：结构体与指针

## 练习 7.1 和 7.2

~~~go
type Product struct {
    Name  string
    Price float64
    Stock int
}

product := Product{
    Name:  "键盘",
    Price: 199,
    Stock: 10,
}

fmt.Println(product.Name, product.Price)
~~~

## 练习 7.3

~~~go
func addStock(product *Product, amount int) {
    product.Stock += amount
}
~~~

调用 addStock(&product, 5)，把 product 的地址传入函数。

## 练习 7.4

~~~go
func (p Product) TotalValue() float64 {
    return p.Price * float64(p.Stock)
}
~~~

这个方法只读取字段，不修改 Product，所以值接收者足够。
