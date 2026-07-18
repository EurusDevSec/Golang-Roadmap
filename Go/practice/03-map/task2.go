package main

import "fmt"

func main() {

	ports := map[string]int{
		"http":  80,
		"https": 443,
	}
	ports["smtp"] = 25
	fmt.Println(ports)

}
