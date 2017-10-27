package main

import "fmt"

func main() {

	// sorting in increase order
	fmt.Println("----------------- Start Ordering Ascending -----------------")
	var a = []int{5, 2, 4, 6, 1, 3}

	var b = a
	asc := sort(b, "asc")
	fmt.Println("\n-- Slice after ascending sort: ", asc, " --\r\n")

	fmt.Println("")

	// sorting in increase order
	fmt.Println("\n----------------- Start Ordering Descending -----------------")
	desc := sort(a, "desc")
	fmt.Println("\n-- Slice after descending sort: ", desc, " --\r\n")

}

func sort(a []int, order string) []int {

	fmt.Println("-- Initial value of array: ", a, " --\r\n")

	for j := 1; j < len(a); j++ {

		// store current value of j
		val := a[j]

		// move one index before j
		i := j - 1

		// work only if i is greater than index 0
		for i > -1 {

			// exit loop based on sorting type
			if order == "asc" && a[i] < val {
				break
			}
			if order == "desc" && a[i] > val {
				break
			}

			// switch right element to left
			a[i+1] = a[i]

			// switch left element to right
			a[i] = val

			// decrement i
			i--
			fmt.Println("Slice inside second for loop: ", a)
		}
	}
	return a
}
