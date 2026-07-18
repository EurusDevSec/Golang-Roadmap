package main

import (
	"fmt"
)

func main() {

	system := []string{"go", "linux", "go", "docker", "go"}
	counts := make(map[string]int)

	for _, value := range system {
		counts[value]++
	}

	for key, value := range counts {
		fmt.Printf("%s=%d\n", key, value)
	}
}
