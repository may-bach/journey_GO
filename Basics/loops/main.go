package main

import "fmt"

func main(){

	//For

	for i:=0; i<10; i++{
		fmt.Println(i)
	}

	//While

	i := 0

	for i<3 {
		fmt.Println(i)
		i++
	}

	//Infinite

	count := 0

	for {
		fmt.Println(count)
		if count == 3{
			break
		}
		count++
	}

	//Continue

	var isfac bool

	for j := 0; j<20 ;j++ {
		isfac = false
		i := 2

		if j==0 || j==1 {
			continue
		}

		for i< (j/2+1) {

			if j%i == 0 {
				isfac = true
				break
			}
			i++
		}
		if isfac {continue}

		fmt.Println(j)
	}
}