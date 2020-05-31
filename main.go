package main

import (
	"fmt"
	"math"
)

func firstTask()  {
	const rate float64 = 52
	var userInput float64
	fmt.Println("please enter the amount in rubles: ")
	fmt.Scanln(&userInput)
	fmt.Println("Sum in a dollar -", userInput / rate, "\n")
}



func secondTask()  {
	var a float64 = 3
	var b float64 = 3
	var c float64 = 2
	var s float64
	var p float64
	var d float64
	s = (a*b)*0.5
	p = a + b + c
	d = math.Sqrt(a) + math.Sqrt(b)
	fmt.Printf("Площадь треугольника: %f\nПериметр треугольника: %f\nГипотенуза тругольника: %f\n\n", s, p, d)
}

func thirdTask()  {
	var sumMoney float64
	var percent float64
	var years = 5
	var result float64
	fmt.Println("Please enter the full amount: ")
	fmt.Scanln(&sumMoney)
	fmt.Println("Please enter annual rate: ")
	fmt.Scanln(&percent)

	for year := 1; year<= years ; year++ {
		result+=(sumMoney * percent * 365)/(365*100)
	}

	fmt.Println("Sum after 5 years", result+sumMoney)
}

func main()  {
	firstTask()
	secondTask()
	thirdTask()
}