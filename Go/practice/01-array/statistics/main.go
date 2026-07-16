package main

import (
	"fmt"
)

func main() {

	statistics := [5]float64{20.5, 35, 80.5, 70, 44}
	min := statistics[0]
	max := statistics[0]
	sum := 0.0
	for _, value := range statistics {
		if value < min {
			min = value
		}

		if value > max {

			max = value
		}
		sum += value
	}
	avg := sum / float64(len(statistics))
	fmt.Printf("min=%.2f\n", min)
	fmt.Printf("max=%.2f\n", max)
	fmt.Printf("avg=%.2f\n", avg)

}
