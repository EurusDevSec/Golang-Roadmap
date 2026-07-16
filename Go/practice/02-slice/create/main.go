package main

import "fmt"

func main() {
	services := []string{"api", "db", "cache"}

	fmt.Println(services)
	fmt.Println("length=", len(services))

}
