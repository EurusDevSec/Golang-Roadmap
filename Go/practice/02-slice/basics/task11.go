package main

import "fmt"

func main() {

	ports := []int{22, 80, 443, 3000, 8080}

	part := ports[1:4]
	fmt.Printf("original: %v\n", ports)
	fmt.Printf("part: %v", part)
	fmt.Println()
	fmt.Println("--------------------")
	part[0] = 81
	fmt.Printf("original: %v\n", ports)
	fmt.Printf("part: %v", part)

	// Ly do original thay doi theo vi slice cua part thuc su chi luu
	// con tro den ports slice cho nen khi thay doi phan tu dau tien
	// cua part thif no cung thay doi o port
}
