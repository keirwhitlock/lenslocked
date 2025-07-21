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
	numbers := []int{1, 2, 3}
	fmt.Println(numbers[4])

	err := Connect()
	if err != nil {
		fmt.Println(err)
	}
	//err := CreateUser()
	//if err != nil {
	//	log.Println(err)
	//}
	//err = CreateOrg()
	//if err != nil {
	//	log.Println(err)
	//}
}
