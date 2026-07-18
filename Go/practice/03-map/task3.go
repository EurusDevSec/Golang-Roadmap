package main

import "fmt"

func main() {

	replicas := map[string]int{
		"api":    2,
		"worker": 1,
	}

	fmt.Println(replicas["api"])

	replicas["api"] = 3
	fmt.Println(replicas["api"])

}
