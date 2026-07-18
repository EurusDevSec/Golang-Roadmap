package main

import "fmt"

func main() {

	defaultPorts := map[string]int{
		"http":     80,
		"https":    443,
		"ssh":      22,
		"postgres": 5432,
	}

	fmt.Println(defaultPorts["https"])
	if value, ok := defaultPorts["redis"]; ok {
		fmt.Println("co key redis,value = ", value)

	} else {
		fmt.Println("khong co key redis")

	}
}
