package main

import (
	"errors"
	"fmt"
)

func Connect() error {
	return errors.New("connection failed")
}

//
//func CreateUser() error {
//	err := Connect()
//	if err != nil {
//		return fmt.Errorf("create user: %w", err)
//	}
//	return nil
//}
//
//func CreateOrg() error {
//	err := CreateUser()
//	if err != nil {
//		return fmt.Errorf("create org: %w", err)
//	}
//	return nil
//}

func main() {
	fib := []int{1, 1, 2, 3, 5, 8}
	Demo(fib...)
}

func Demo(numbers ...int) {
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func Sum(numbers ...int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
