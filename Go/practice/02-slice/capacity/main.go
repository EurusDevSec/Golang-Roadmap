package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 90, 70, 60}
	slice := numbers[:3]
	fmt.Println(cap(slice))
	newSlice := append(slice, 100, 200)
	fmt.Println(cap(newSlice))
}
