package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		s := 0.5*a*t*t + v0*t + s0
		return s
	}
}

func GetValueFromUser(key string) float64 {
	var input string
	fmt.Printf("\nEnter %v:", key)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Printf("Invalid input !")
	}
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("Invalid input !")
	}
	return value
}

func main() {
	var valueMap map[string]float64 = map[string]float64{"Acceleration": 0, "Initial Velocity": 0, "Initial Displacement": 0, "Time": 0}

	for key, _ := range valueMap {
		if key != "Time" {
			valueMap[key] = GetValueFromUser(key)
		}
	}
	fn := GenDisplaceFn(valueMap["Acceleration"], valueMap["Initial Velocity"], valueMap["Initial Displacement"])
	valueMap["Time"] = GetValueFromUser("Time")

	fmt.Printf("displacement = %v", fn(valueMap["Time"]))
}
