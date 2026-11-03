package main

import (
	"errors"
	"fmt"
)

var ErrInvalidAge = errors.New("年龄不能小于 0")

func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("校验用户资料失败: %w", ErrInvalidAge)
	}
	return nil
}

func main() {
	err := validateAge(-1)
	if err != nil {
		fmt.Println("发生错误：", err)
		if errors.Is(err, ErrInvalidAge) {
			fmt.Println("原因：年龄不合法")
		}
		return
	}

	fmt.Println("年龄合法")
}
