package main

import "fmt"

func main() {
	days := []int{13, 12, 15, 11, 9, 12, 16}
	resp := whemWarmer_V0(days)
	fmt.Println(resp)
}

func whemWarmer_V0(days []int) []int {
	arr := make([]int, len(days))

	for i := 0; i < len(days); i++ {
		for j := i + 1; j < len(days); j++ {
			if days[i] < days[j] {
				arr[i] = j - i
				break
			}
		}
	}
	return arr
}
