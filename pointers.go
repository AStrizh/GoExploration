package main

import "fmt"

func PointerTest(){
	age := 32

	agePointer := &age

	fmt.Println(*agePointer)
}