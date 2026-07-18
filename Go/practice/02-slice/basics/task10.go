package main

import "fmt"

func main() {

	services := []string{"api"}
	more := []string{"db", "cache", "worker"}

	services = append(services, more...)
	fmt.Print(services)
}
