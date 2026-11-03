package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	results := make(chan int, 3)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(number int) {
			defer wg.Done()
			time.Sleep(time.Duration(number) * 100 * time.Millisecond)
			results <- number * number
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("收到结果：", result)
	}

	fmt.Println("所有任务完成")
}
