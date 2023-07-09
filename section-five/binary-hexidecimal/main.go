package main

import "fmt"

func main() {
	numberToFormat := 42

	fmt.Printf("42 in binary is, %b \n", numberToFormat)
	fmt.Printf("42 in hexidecimal is, %x \n", numberToFormat)

	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("The binary representation of %d is %b \n", a, a)
	fmt.Printf("The binary representation of %d is %b \n", b, b)
	fmt.Printf("The binary representation of %d is %b \n", c, c)
	fmt.Printf("The binary representation of %d is %b \n", d, d)
	fmt.Printf("The binary representation of %d is %b \n", e, e)
	fmt.Printf("The binary representation of %d is %b \n", f, f)

	fmt.Printf("The hexidecmial representation of %d is %#x \n", a, a)
	fmt.Printf("The hexidecmial representation of %d is %#x \n", b, b)
	fmt.Printf("The hexidecmial representation of %d is %#x \n", c, c)
	fmt.Printf("The hexidecmial representation of %d is %#x \n", d, d)
	fmt.Printf("The hexidecmial representation of %d is %#x \n", e, e)
	fmt.Printf("The hexidecmial representation of %d is %#x \n", f, f)

}
