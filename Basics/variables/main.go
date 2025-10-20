package main

import "fmt"

func main() {
	var name string // declares a var named "name" of type string
	name = "Izumi"

	iq := 67.41 // := assumes the datatype of the value passed

	fmt.Println("Name: ", name)
	fmt.Println("Average iq of a ferrari engineer: ", iq)

	var retarded bool
	retarded = true

	var def_int int // No need to specify value = 0 like C

	_ = def_int //blank identifier to tell GO this var isn't used yet (not necessary, only to remove "unsused" warning)

	const pi = 3.141593

	fmt.Println("Am i retarded? ", retarded)

	fmt.Printf("I'm %s, im smarter than an avg ferrari engineer who has an iq of %f. \nIntegers are intialzed with value %d. Area of a circle is %.2f x r x r", name, iq, def_int, pi)

}
