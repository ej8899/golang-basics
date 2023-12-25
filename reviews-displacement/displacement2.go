package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func ReadUserImput() (float64, float64, float64, float64) {
	var acceleration float64
	var initialVelocity float64
	var initialDisplacement float64
	var time float64

	scanner := bufio.NewScanner(os.Stdin)

	log.Print("Please enter acceleration:")
	for scanner.Scan() {
		acceleration, _ = strconv.ParseFloat(scanner.Text(), 64)
		break
	}

	log.Print("Please enter initial velocity:")
	for scanner.Scan() {
		initialVelocity, _ = strconv.ParseFloat(scanner.Text(), 64)
		break
	}

	log.Print("Please enter initial displacement:")
	for scanner.Scan() {
		initialDisplacement, _ = strconv.ParseFloat(scanner.Text(), 64)
		break
	}

	log.Print("Please enter time:")
	for scanner.Scan() {
		time, _ = strconv.ParseFloat(scanner.Text(), 64)
		break
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return acceleration, initialVelocity, initialDisplacement, time
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*(math.Pow(t, 2)*0.5) + (v0 * t) + (s0))
	}
	return fn
}

func main() {
	acceleration, initialVelocity, initialDisplacement, time := ReadUserImput()
	fmt.Printf("Acceleration:         %f\n", acceleration)
	fmt.Printf("Initial velocity:     %f\n", initialVelocity)
	fmt.Printf("Initial displacement: %f\n", initialDisplacement)
	fmt.Printf("Time:                 %f\n", time)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Printf("Displacement:         %f\n", fn(time))
}