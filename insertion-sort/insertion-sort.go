package main

import "fmt"

func main() {

	// sorting in increase order
	fmt.Println("----------------- Start Ordering Ascending -----------------")
	var a = [6]int{5, 2, 4, 6, 1, 3}

	for j := 1; j < len(a); j++ {
		val := a[j]
		fmt.Println("Index j: ", j)
		fmt.Println("Value: ", val)
		i := j - 1

		fmt.Println("Index i: ", i)
		for i > -1 && a[i] > val {
			fmt.Println("Index i inside second loop: ", i)
			a[i+1] = a[i]
			i = i - 1
			a[i+1] = val
		}
		fmt.Println("-- Value of a: ", a, " --\r\n")
	}

	fmt.Println("")

	// sorting in increase order
	fmt.Println("----------------- Start Ordering Descending -----------------")
	var b = [6]int{5, 2, 4, 6, 1, 3}

	for j := 1; j < len(b); j++ {
		val := a[j]
		fmt.Println("Index j: ", j)
		fmt.Println("Value: ", val)
		i := j - 1

		fmt.Println("Index i: ", i)
		for i > -1 && b[i] < val {
			fmt.Println("Index i inside second loop: ", i)
			b[i+1] = b[i]
			i = i - 1
			b[i+1] = val
		}
		fmt.Println("-- Value of a: ", b, " --\r\n")
	}
}
