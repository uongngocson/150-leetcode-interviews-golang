package main
import "fmt"

func candy(ratings []int) int {
	n := len(ratings)

	candies := make([]int, n)

	// Mỗi đứa ít nhất 1 viên
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	// Left -> Right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Right -> Left
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	total := 0
	for _, c := range candies {
		total += c
	}

	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ratings := []int{1,0,2}
	fmt.Println(candy(ratings))
}