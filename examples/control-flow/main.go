package main

import "fmt"

func main() {
	score := 82

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("循环数字：", i)
	}

	day := 2
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	default:
		fmt.Println("其他日期")
	}
}
