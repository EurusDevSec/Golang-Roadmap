package main

import "fmt"

func double(values []int) {
	//TODO
	for index := range values {
		values[index] *= 2
	}

}

func addService(services []string, name string) []string {
	// TODO

	services = append(services, name)
	return services
}
func main() {
	// nums := []int{1, 2, 3}
	// double(nums)
	// fmt.Println(nums)

	services := []string{"api", "db"}
	services = addService(services, "cache")
	fmt.Println(services)
}
