package main

import "fmt"

func addThree_Copy(score int) {
	score += 3
	fmt.Printf("Inside addThree_Copy (the copy): %d\n", score)
}

func addThree_Pointer(scorePtr *int) {
	*scorePtr += 3
	fmt.Printf("Inside addThree_Pointer (the original): %d\n", *scorePtr)
}

func main() {
	gameScore := 100
	fmt.Printf("Score: %d\n", gameScore)

	fmt.Println("CURRY FROM WAY DOWNTOWN !")
	fmt.Println("Trying to change the score:")
	addThree_Copy(gameScore)
	fmt.Printf("New score: %d\n", gameScore)       // no change because... its basic pointers man you should know ts

	fmt.Println("--------------------")

	fmt.Println("Changing the original via pointer:")
	addThree_Pointer(&gameScore)
	fmt.Printf("New score: %d\n", gameScore)
}