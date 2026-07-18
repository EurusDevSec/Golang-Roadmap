package main

import "fmt"

func main() {

	services := make([]string, 3)

	services[0] = "api"
	services[1] = "db"
	services[2] = "cache"

	fmt.Println("services=", services)
	fmt.Printf("len=%d\ncap=%d\n", len(services), cap(services))

}
