package main

import "fmt"

func checkIfEven() {
	// проверка чётности числа вводимого пользователем
	var userNumber int
	fmt.Println("Введите чиисло: ")
	fmt.Scanln(&userNumber)
	if userNumber%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}

func checkIfDivisibleOnThree() {
	// проверка делимости числа на 3 без остатка вводимого пользователем
	var userNumber int
	fmt.Println("Введите число: ")
	fmt.Scanln(&userNumber)
	if userNumber%3 == 0 {
		fmt.Println("Делится на 3 без остатка")
	} else {
		fmt.Println("Не делится на 3 без остатка")
	}
}

func fibo() {
	// последовательность фебоначи
	for i, j := 0, 1; j < 100; i, j = i+j, i {
		if i < 100 {
			fmt.Println(i)
		}
	}
}
