package main

import (
	"fmt"
)

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println(numbers)

	//task 2, Get third element, last element
	fmt.Println("first=", numbers[0])
	fmt.Println("third=", numbers[2])
	fmt.Println("last=", numbers[len(numbers)-1])

	// task 3: change element 20 to element 99

	numbers2 := [3]int{10, 20, 30}
	fmt.Println("before: ", numbers2)
	numbers2[1] = 99
	fmt.Println("after: ", numbers2)
	//task 4 Traversal array

	fmt.Println("------------------")

	system := [4]string{"api", "db", "cache", "worker"}
	for index, value := range system {
		fmt.Println(index, ":", value)
	}

	//task 5 sum
	num3 := [5]int{2, 4, 6, 8, 10}
	sum := 0

	for _, value := range num3 {
		sum += value
	}
	fmt.Println("sum=", sum)

	//task 6
	num6 := [6]int{2, 8, 1, 9, 5, 7}
	count := 0

	for i := 0; i < len(num6); i++ {
		if num6[i] > 5 {
			count++
		}
	}
	fmt.Println("count=", count)

	//task 7
	num7 := [5]int{4, 10, 3, 8, 6}
	max := num7[0]
	for _, value := range num7 {
		if value > max {
			max = value
		}
	}
	fmt.Println("max=", max)
}
