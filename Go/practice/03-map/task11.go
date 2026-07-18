package main

import "fmt"

func main() {

	allowIPs := []string{"10.0.0.1", "10.0.0.2", "10.0.0.1"}
	checkIP := "10.0.0.2"

	whiteList := make(map[string]bool)
	for _, ip := range allowIPs {
		whiteList[ip] = true
	}
	fmt.Println(whiteList)
	allowed := whiteList[checkIP]
	uniqueIps := len(whiteList)

	fmt.Printf("allowed=%v\n", allowed)
	fmt.Printf("unique_ips=%d\n", uniqueIps)
}
