package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"time"
)

func main(){

	fmt.Println("--------F1 Points Tracker--------")

	teamPoints := make(map[string]int)

	var cmdHis []int

	var valid bool = true

	in := bufio.NewReader(os.Stdin)


	for valid{
		
		fmt.Println("What do you wish to do? \n1.Add\n2.Check\n3.Print\n4.History\n5.Quit")
		fmt.Println("---------------------------------")
		fmt.Printf("Choice: ")

		choice,_ := in.ReadString('\n')
		trim := strings.TrimSpace(choice)
		n,_ := strconv.Atoi(trim)

		fmt.Printf("\n")

		if(n!=1 && n!=2 && n!=3 && n!=4 && n!=5){
			fmt.Println("Invalid choice")
			continue
		}

		cmdHis = append(cmdHis,n)

		switch(n){

		case 1:

			fmt.Printf("Enter team name: ")
			
			name,_ := in.ReadString('\n')
			teamName := strings.TrimSpace(name)

			fmt.Printf("\n")
			
			fmt.Printf("Enter points to be added: ")
			input,_ := in.ReadString('\n')
			trimm := strings.TrimSpace(input)
			points,err := strconv.Atoi(trimm)

			fmt.Printf("\n")

			if err != nil {
				fmt.Println("Invalid points")
				fmt.Println("---------------------------------")
				time.Sleep(2 * time.Second)
				continue
			}

			teamPoints[teamName] += points
			
			fmt.Printf("%d points added to %s.\n",points,teamName)
			fmt.Println("---------------------------------")
			time.Sleep(2 * time.Second)
		
		case 2:

			fmt.Printf("Enter team name: ")
			name,_ := in.ReadString('\n')
			teamName := strings.TrimSpace(name)

			fmt.Printf("\n")

			points, ok := teamPoints[teamName]

			if ok{
				fmt.Printf("Points held by %s: %d\n",teamName,points)
				fmt.Println("---------------------------------")
			} else{
				fmt.Printf("No team named %s found.\n",teamName)
				fmt.Println("---------------------------------")
			}
			time.Sleep(2 * time.Second)

		case 3:

			fmt.Println("--------Current Standings--------")
			for name,points := range teamPoints{
				fmt.Printf("%-10s | %d\n",name,points)
			}

			fmt.Println("---------------------------------")
			time.Sleep(2 * time.Second)

		case 4:
			
			fmt.Println("--------History--------")
			for i:=0;i<len(cmdHis);i++ {
				fmt.Printf("%d. ",cmdHis[i])
				fmt.Printf("\n")
			}

			fmt.Println("---------------------------------")
			time.Sleep(2 * time.Second)

		case 5:
			valid = false
		}
	}
}