package main

import "fmt"

func double(values []int) {
	for i := range values {
		values[i] *= 2
	}
}

func main() {
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)
	double(numbers)

	cloned := append([]int(nil), numbers...)
	cloned[0] = 100

	fmt.Println("original:", numbers)
	fmt.Println("clone:", cloned)
}
