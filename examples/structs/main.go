package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) Birthday() {
	u.Age++
}

func main() {
	user := User{Name: "小明", Age: 20}
	fmt.Printf("用户：%s，年龄：%d\n", user.Name, user.Age)

	user.Birthday()
	fmt.Printf("生日后年龄：%d\n", user.Age)

	pointer := &user
	pointer.Name = "小明（已更新）"
	fmt.Println(pointer.Name)
}
