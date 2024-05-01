package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Hello World\n")

	fmt.Println(investmentCalculator(1000,5.5,10))

	y := 7
	x := 5


	fmt.Println(x+y)


}

func investmentCalculator(amount float64, rate float64, years int) float64{
	var futureValue = amount * math.Pow(1+rate/100,float64(years))
	return futureValue
}