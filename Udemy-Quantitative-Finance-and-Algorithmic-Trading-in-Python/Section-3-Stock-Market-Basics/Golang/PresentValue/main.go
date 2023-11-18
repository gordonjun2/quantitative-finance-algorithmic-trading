package main

import (
	"fmt"
	"math"
)

func main() {

	x := 100.0
	r := 0.05
	n := 5.0

	fmt.Printf("Future value (discrete model) of x: %f\n", future_discrete_value(x, r, n))
	fmt.Printf("Present value (discrete model) of x: %f\n", present_discrete_value(x, r, n))
	fmt.Printf("Future value (continuous model) of x: %f\n", future_continuous_value(x, r, n))
	fmt.Printf("Present values (continuous model) of x: %f\n", present_continuous_value(x, r, n))
}

func future_discrete_value(x, r, n float64) float64 {
	return x * math.Pow(1+r, n)
}

func present_discrete_value(x, r, n float64) float64 {
	return x * math.Pow(1+r, -n)
}

func future_continuous_value(x, r, t float64) float64 {
	return x * math.Exp(r*t)
}

func present_continuous_value(x, r, t float64) float64 {
	return x * math.Exp(-r*t)
}
