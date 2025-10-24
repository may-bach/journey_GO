package main

import "fmt"

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println("original primes: ", primes)

	var slice []int = primes[1:6] // no number for size while slicing

	slice[0] = 67

	fmt.Println("Slice: ", slice)
	fmt.Println("better primes: ", primes) //any change in the slice of an array is reflected in the original array, similar to pointer
}
