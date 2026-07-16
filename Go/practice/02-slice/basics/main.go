package main

import (
	"fmt"
)

func main() {

	services := []string{"api", "db", "cache"}
	servicesWithWorker := append(services, "worker")
	fmt.Println("before: ", services)
	fmt.Println("after: ", servicesWithWorker)

	fmt.Println("-----------------------------------")
	services[1] = "postgres"
	fmt.Println(services)
	fmt.Println("---------------------------------")

	sum := 0

	numbersForSum := []int{5, 10, 15, 20}
	for _, value := range numbersForSum {
		sum += value
	}
	fmt.Println("sum=", sum)

	fmt.Println("-------------------------------")
	numbers := []int{1, 2, 3, 4, 5, 6}
	evenNumbers := make([]int, 0)

	for _, value := range numbers {
		if value%2 == 0 {
			evenNumbers = append(evenNumbers, value)
		}
	}
	fmt.Println(evenNumbers)

	fmt.Println("--------------------")
	// findNameService := []string{"api", "db", "cache"}
	// var target string
	// fmt.Scan(&target)
	// found := false
	// for _, value := range findNameService {
	// 	if value == target {
	// 		found = true
	// 	}
	// }
	// fmt.Println("found=", found)

	fmt.Println("-----------------------")
	array := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}
	slice = append(slice, 4)

	fmt.Println("array=", array, "length=", len(array))
	fmt.Println("slice=", slice, "length=", len(slice))
}
