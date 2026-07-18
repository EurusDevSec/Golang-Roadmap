package main

import (
	"fmt"
)

func main() {

	ports := map[string]int{
		"http":     80,
		"https":    443,
		"ftp":      21,
		"ssh":      22,
		"mysql":    3306,
		"postgres": 5432,
	}

	for key, value := range ports {
		fmt.Printf("%s=%d\n", key, value)
	}

}
