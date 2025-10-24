package main

import "fmt"

func main() {

	var points [6]int

	points = [6]int{4, 8, 2, 12, 0, 28} //the size is necessary

	fmt.Println("Points scored by Leclerc in the last 6 races: ", points)

	points[4] = 10

	fmt.Println("Points he would've scored in Dutch Gp if Kimi wasn't a terrorist: ", points)

	vowels := [...]string{"a", "e", "i", "o", "u"} //... makes the compiler infer the size

	fmt.Println("James ", vowels, " is the best team principal")

	fmt.Printf("There are %d vowels in the alphabet\n", len(vowels))

}
