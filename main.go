package main

import (
	"fmt"
)

func main() {
	var i int = 0
	for i < 5 {
		fmt.Printf("Nilai i = %d\n", i)
		i++
		if i == 5 {
			for j := 0; j <= 10; j++ {
				if j == 5 {
					input := "САШАРВО"
					for index, value := range input {
						fmt.Printf("character %U '%c' starts at byte position %d\n", value, value, index)

					}
				} else {
					fmt.Printf("nilai J = %d\n", j)
				}
			}
		}
	}
}
