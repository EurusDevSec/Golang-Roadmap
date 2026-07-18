package main

import "fmt"

func main() {

	envs := []string{"prod", "staging", "prod", "dev", "prod", "staging"}

	counts := make(map[string]int)

	for _, value := range envs {
		counts[value]++
	}
	for key, value := range counts {
		fmt.Println(key, "=", value)
	}
}
