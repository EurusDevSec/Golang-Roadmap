package main

import "fmt"

func main() {

	services := make([]string, 0, 3)
	fmt.Printf("before: [] len=%d cap=%d\n", len(services), cap(services))
	services = append(services, "api")
	services = append(services, "db")
	services = append(services, "cache")
	fmt.Printf("after: [] len=%d cap=%d\n", len(services), cap(services))

}
