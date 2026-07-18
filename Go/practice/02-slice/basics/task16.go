package main

import "fmt"

func main() {

	services := []string{"api", "db", "cache", "worker"}
	index := 1

	services = append(services[:index], services[index+1:]...)
	fmt.Println(services)
}
