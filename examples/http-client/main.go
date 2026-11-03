package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://httpbin.org/get",
		nil,
	)
	if err != nil {
		panic(err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		panic(err)
	}

	fmt.Println("HTTP 状态：", resp.Status)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println("请求没有成功")
		return
	}

	fmt.Println(string(body))
}
