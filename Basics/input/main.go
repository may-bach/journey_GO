package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main(){

	fmt.Printf("Enter your identification: ")

	reader := bufio.NewReader(os.Stdin)

	id, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error acquiring identity: ", err)
		return
	}

	id = strings.TrimSpace(id)

	fmt.Printf("Authentication successful. Welcome Spirit user, %s\n", id)

}