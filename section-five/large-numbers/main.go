package main

import (
	"fmt"
	"math"
)

func main() {
	var signed int8 = math.MaxInt8
	var unSigned uint8 = math.MaxUint8

	fmt.Printf("The largest number possible for type %T is %d\n", signed, signed)
	fmt.Printf("The largest number possible for type %T is %d\n", unSigned, unSigned)
}
