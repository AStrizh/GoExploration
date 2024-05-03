package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Print("Hello World\n")

	// var x int
	// y := 7
	// fmt.Scan(&x)

	// var result = x + y

	// if result < 10 {
	// 	fmt.Println("Too Small!")
	// } else {
	// 	fmt.Println("That works!")
	// }

	// for i := x; i > 0; i-- {
	// 	fmt.Println(i)
	// }

	var number = 12
	os.WriteFile("output.txt", []byte(strconv.Itoa(number)+"\n"), 0644)

	investmentCalculator(1000, 2.5, 30)

	fmt.Println(randomdata.ProvinceForCountry("FR"))

}
