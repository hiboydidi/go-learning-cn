package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.CreateTemp("", "go-learning-*.txt")
	if err != nil {
		panic(err)
	}

	path := file.Name()
	defer os.Remove(path)

	if _, err := file.WriteString("第一行：学习 Go\n第二行：练习文件读写\n"); err != nil {
		file.Close()
		panic(err)
	}

	if err := file.Close(); err != nil {
		panic(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
