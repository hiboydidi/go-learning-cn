package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println("2 + 3 =", add(2, 3))
	fmt.Println("sum =", sum(1, 2, 3, 4))

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("计算失败：", err)
		return
	}
	fmt.Println("10 / 2 =", result)
}
