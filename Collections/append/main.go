package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	points := make([]int, 3, 5) //length is 3, and max capacity 5

	input := bufio.NewReader(os.Stdin)

	for i := 0; i < 3; i++ {

		point, err := input.ReadString('\n')

		if err != nil {
			fmt.Printf("Error reading point: %v\n", err)
			return
		}

		trim := strings.TrimSpace(point)
		n, err := strconv.Atoi(trim) // converting string input to integer  (why tf is ts so complicated)

		if err != nil {
			fmt.Printf("Invalid point: %v\n", err)
			return
		}

		points[i] = n
	}

	fmt.Println("Max points: ", points)
	fmt.Println("Number of races: ", len(points)) //lenth of the array
	fmt.Println("Total races: ", cap(points))     //max length of array

	points = append(points, 25)
	points = append(points, 18)

	fmt.Println("Max points: ", points)
	fmt.Println("Number of races: ", len(points))
	fmt.Println("Total races: ", cap(points))

	points = append(points, 20)
	points = append(points, 2)

	fmt.Println("Max points: ", points)
	fmt.Println("Number of races: ", len(points)) //length of the array
	fmt.Println("Total races: ", cap(points))     //max length of the array automatically doubled when more data than max capacity is appended

	points = append(points, 10)
	points = append(points, 67)

	fmt.Println("Max points: ", points)
	fmt.Println("Number of races: ", len(points))
	fmt.Println("Total races: ", cap(points))

}
