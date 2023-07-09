package main

import "fmt"

func main() {
	num1, num2, num3 := 747, 911, 90210

	fmt.Printf("%d\t %b\t\t %#X\n", num1, num1, num1)
	fmt.Printf("%d\t %b\t\t %#X\n", num2, num2, num2)
	fmt.Printf("%d\t %b\t %#X\n", num3, num3, num3)
}
