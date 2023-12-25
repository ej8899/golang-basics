package main

import ("fmt")

// generate displacement based on supplied forumula
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5*a*t*t) + (vo*t) + so
	}
}

func main() {
	var acceleration, initialVelocity, initialDisplacement float64

	fmt.Print("Enter acceleration: ")
	fmt.Scanln(&acceleration)

	fmt.Print("Enter initial velocity: ")
	fmt.Scanln(&initialVelocity)

	fmt.Print("Enter initial displacement: ")
	fmt.Scanln(&initialDisplacement)

	// Generate displacement function
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	// Get time for displacement calc
	var time float64
	fmt.Print("Enter time: ")
	fmt.Scanln(&time)

	// Compute and print the displacement after time
	displacement := fn(time)
	fmt.Println("Displacement after", time, "seconds:", displacement)
}
