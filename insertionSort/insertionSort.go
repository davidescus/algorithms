package main

import "fmt"

type Sort []int

func main() {

	var a = []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Original slice: ", a)

	s := Sort(a)
	fmt.Println("Ascending sort: ", s.sortAsc())

	fmt.Println("Descendingsort: ", s.sortDesc())
}

func (s Sort) sortAsc() []int {
	return sort(s, "asc")
}

func (s Sort) sortDesc() []int {
	return sort(s, "desc")
}

func sort(a []int, order string) []int {

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
			// decrement i for next iteration
			i--
			//fmt.Println("Slice inside second for loop: ", a)
		}
	}
	return a
}
