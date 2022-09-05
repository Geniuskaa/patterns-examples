package main

import (
	"fmt"
	"time"
)

//	Разбор задачи: https://www.youtube.com/watch?v=GhiRlhPlJ9Q

func main() {
	start := time.Now()
	//count := PathsCount_V0(17, 19) // VERY SLOW
	count := PathsCount_V1(17, 19) // FASTEST
	finish := time.Since(start).Microseconds()
	fmt.Println("Worked: ", finish)
	fmt.Println(count)
}

func PathsCount_V0(n int, m int) int {
	if n < 1 || m < 1 {
		return 0
	}

	if n == 1 && m == 1 {
		return 1
	}

	return PathsCount_V0(n-1, m) + PathsCount_V0(n, m-1)
}

// Fastest variant
func PathsCount_V1(n int, m int) int {
	arr := make([][]int, n+1)
	for i := range arr {
		arr[i] = make([]int, m+1)
	}
	return helper(n, m, arr)
}

func helper(n int, m int, arr [][]int) int {
	if n < 1 || m < 1 {
		return 0
	}

	if n == 1 && m == 1 {
		return 1
	}

	if arr[n][m] != 0 {
		return arr[n][m]
	}
	arr[n][m] = helper(n-1, m, arr) + helper(n, m-1, arr)
	return arr[n][m]
}
