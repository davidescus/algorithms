package main

import "fmt"

// signature of compare function
type compare func(int, int) bool

func asc(v1, v2 int) bool {
	return v1 < v2
}

func desc(v1, v2 int) bool {
	return v1 > v2
}

func main() {

	var a = []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Original slice: ", a)

	fmt.Println("Ascending sort: ", sort(a, asc))
	fmt.Println("Descending sort: ", sort(a, desc))
}

func sort(a []int, cmp compare) []int {

	for j := 1; j < len(a); j++ {
		// store current value of j
		val := a[j]
		// move one index before j
		i := j - 1

		// work only if i is greater than index 0
		for i > -1 {
			// exit loop based on sorting type
			if cmp(a[i], val) {
				break
			}
			// switch right element to left
			a[i+1] = a[i]
			// switch left element to right
			a[i] = val
			// decrement i for next iteration
			i--
			//fmt.Println("Slice inside second for loop: ", a)
		}
	}
	return a
}
