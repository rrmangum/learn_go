package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

var packageScope = 123

const meaningOfLife = 42

func main() {
	blockScope := 123
	blockScope2 := 456

	fmt.Println(packageScope)
	fmt.Println(meaningOfLife)

	fmt.Printf("The value of blockScope is %v and the type is %T\n", blockScope, blockScope)
	fmt.Println(blockScope2)

	fmt.Println(puppy.Bark())
}
