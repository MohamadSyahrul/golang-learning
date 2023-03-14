package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Selamat Malam"
	kata := strings.Split(input, "")

	for _, k := range kata {
		fmt.Println(k)
	}

	hitung := make(map[rune]int)
	for _, r := range input {
		hitung[r]++
	}

	hitungString := make(map[string]int)
	for key, nilai := range hitung {
		hitungString[string(key)] = nilai
	}

	fmt.Println(hitungString)
}
