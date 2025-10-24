package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Enter your identification: ")

	reader := bufio.NewReader(os.Stdin) // new input scanner variable

	id, err := reader.ReadString('\n') //defining err in case of invalid variable passed etc...

	if err != nil {
		fmt.Printf("Error acquiring identity: ", err)
		return
	}

	id = strings.TrimSpace(id) // removing spaces

	fmt.Printf("Authentication successful. Welcome Spirit user, %s\n", id)

}
