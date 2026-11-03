package main

import "fmt"

func main() {
	scores := []int{80, 90, 100}
	scores = append(scores, 95)

	total := 0
	for _, score := range scores {
		total += score
	}
	fmt.Println("成绩：", scores)
	fmt.Println("平均分：", float64(total)/float64(len(scores)))

	counts := map[string]int{}
	words := []string{"go", "python", "go", "java", "go"}
	for _, word := range words {
		counts[word]++
	}
	fmt.Println("go 出现次数：", counts["go"])
}
