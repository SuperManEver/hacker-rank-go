package main

import (
	"fmt"
	"strings"
)

// func timeConversion(s string) string {
// 	// Write your code here

// }

func main() {
	// fmt.Println("Helo world")

	input := "07:05:45PM"
	values := strings.Split(input, ":")

	fmt.Println(values)
}
