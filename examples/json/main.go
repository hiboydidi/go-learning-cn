package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func main() {
	user := User{Name: "小明", Age: 20, Active: true}

	data, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	var decoded User
	if err := json.Unmarshal(data, &decoded); err != nil {
		panic(err)
	}
	fmt.Printf("解码结果：%+v\n", decoded)
}
