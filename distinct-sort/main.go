package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{1, 8, 4, 2, 9, 3, 2, 1, 8, 4, 9, 12, 2, 56, 6}
	sort.Ints(a)

	s := []string{"A", "B", "L", "A", "L", "P", "W", "B", "A", "S", "Z", "Y", "T", "O", "P"}
	sort.Strings(s)

	b1 := distinct(a)
	b2 := distinct(s)

	fmt.Println(b1)
	fmt.Println(b2)
}

// distinct accepts only sorted slices
func distinct[T comparable](arr []T) []T {
	b := make([]T, 0, len(arr))

	previousIndex := -1
	prevCutIndex := 0
	sameValue := false
	sameCounter := 0
	for i, v := range arr {
		if previousIndex == -1 {
			previousIndex = i
			continue
		}

		if v != arr[previousIndex] {
			if sameValue {
				if len(b) == 0 {
					b = append(b, arr[:i-sameCounter]...)
					sameValue = false
					prevCutIndex = i
					previousIndex = i
					sameCounter = 0
				} else {
					b = append(b, arr[prevCutIndex:i-sameCounter]...)
					sameValue = false
					prevCutIndex = i
					previousIndex = i
					sameCounter = 0
				}
			} else {
				previousIndex = i
				if i+1 == len(arr) {
					b = append(b, arr[prevCutIndex:]...)
				}
			}
		} else {
			sameValue = true
			sameCounter++
			if i+1 == len(arr) {
				b = append(b, arr[prevCutIndex:i-sameCounter+1]...)
			}

		}
	}
	return b
}
