package main

import (
	"fmt"
)

func main() {
	var isPalindrome = func(str string) map[string]int {
		var data = map[string]int{}

		for i := range str {
			var x string = string(byte(str[i]))
			data[x] = 0

			fmt.Println(x)

		}

		for i := range str {
			var x string = string(byte(str[i]))
			countCharacter(data, x)
		}

		return data

	}("selamat malam")

	fmt.Println(isPalindrome)
}

func countCharacter(c map[string]int, x string) {
	for key := range c {
		if key == x {
			c[key] += 1
		}
	}
}
