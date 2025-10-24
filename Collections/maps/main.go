package main

import "fmt"

func main(){

	driverAges :=  make(map[string]int)
	driverAges["izumi"] = 18
	driverAges["charles"] = 27
	driverAges["alex"] = 28
	driverAges["lando"] = 25

	fmt.Println("Izumi's age: ",driverAges["izumi"])

	teamPoints := map[string]int{
		"mercedes" : 210,
		"ferrari" : 207,
		"redbull" : 243,
		"mclaren" : 527,
		"haas" : 27,
	}

	fmt.Println("Team standings: ", teamPoints) //not very nice, prolly wont use

	delete(teamPoints, "haas")
	teamPoints["audi"] = 0

	points, ok := teamPoints["haas"]
	if ok {
		fmt.Println("haas points: ", points)
	} else{
		fmt.Println("Imagine getting replaced by audi") //lmao
	}

	for key,value := range driverAges {
		fmt.Printf("Driver: %s, Age: %d\n", key, value) //looks better
	}

}