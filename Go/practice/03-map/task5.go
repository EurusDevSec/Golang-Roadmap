package main

import "fmt"

func main() {

	services := map[string]string{"api": "running", "db": "stopped"}
	delete(services, "db")
	fmt.Println(services)
}
