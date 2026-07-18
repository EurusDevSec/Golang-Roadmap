package main

import "fmt"

func main() {

	services := map[string]string{
		"api": "running", "db": "stopped",
	}
	// value, ok := services["api"]
	// fmt.Println(value, ok)
	value, ok := services["cache"]
	fmt.Println(value, ok)

}
