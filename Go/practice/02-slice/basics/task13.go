package main

import "fmt"

func main() {
	original := []string{"dev", "staging", "prod"}
	cloned := make([]string, len(original))
	copy(cloned, original)
	cloned[0] = "local"
	fmt.Printf("original: %v\n", original)
	fmt.Printf("cloned: %v\n", cloned)

}
