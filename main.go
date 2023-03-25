package main

import (
	"errors"
	"fmt"
)

func main() {
	defer catchErr()

	var password string

	fmt.Scanln(&password)
	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}

}

func validPassword(password string) (string, error) {
	panjang := len(password)
	if panjang < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}
	return "valid password", nil
}
