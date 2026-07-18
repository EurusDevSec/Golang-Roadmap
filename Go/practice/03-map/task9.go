package main

import "fmt"

func main() {

	services := map[string]string{
		"api": "running",
		"db":  "stopped",
		"mq":  "running",
	}
	fmt.Println("before=", services["db"])
	services["db"] = "running"
	fmt.Println("after=", services["db"])

}
