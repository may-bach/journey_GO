package main

import "fmt"
import "os"
import "strings"
import "bufio"
import "strconv"

func square(n1 int) int {         
	sum := n1*n1
	return sum
}

func playerName()(fname string,lname string) {
	fname = "Michael"
	lname = "Jordan"
	return
}

func playerInfo() (string,int){
	return "Steph", 30
}

func main(){

	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter number: ")
	n,_ := in.ReadString('\n')
	trim := strings.TrimSpace(n)

	num, err := strconv.Atoi(trim)

	if err!=nil {
		fmt.Println("Invalid number")
	}

	square := square(num)

	fmt.Println("Square of number: ",square)

	name, number := playerInfo()

	// name, _ :=playerInfo()   for just name

	fmt.Printf("%s's number is %d\n",name,number)

	fname,lname := playerName()

	fmt.Printf("Player Name: %s %s\n",fname,lname)

}