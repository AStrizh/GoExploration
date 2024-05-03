package main

import "math"

func investmentCalculator(amount float64, rate float64, years int) (futureValue float64) {
	futureValue = amount * math.Pow(1+rate/100, float64(years))
	//return futureValue
	return
}