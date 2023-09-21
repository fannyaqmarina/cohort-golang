package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FizzBuzz(10))
}

func FizzBuzz(n int) []string {
	var result []string
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			str := strconv.Itoa(i)
			result = append(result, str)
		}
	}
	return result

}
