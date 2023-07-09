package main

import "fmt"

type ByteSize int

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	// ZB
	// YB
)

func main() {
	fmt.Printf("Decimal: %d \t\t\t Binary: %b\n", KB, KB)
	fmt.Printf("Decimal: %d \t\t Binary: %b\n", MB, MB)
	fmt.Printf("Decimal: %d \t\t Binary: %b\n", GB, GB)
	fmt.Printf("Decimal: %d \t\t Binary: %b\n", TB, TB)
	fmt.Printf("Decimal: %d \t Binary: %b\n", PB, PB)
	fmt.Printf("Decimal: %d \t Binary: %b\n", EB, EB)
}

/*
PB		    1125899906842624
EB		 1152921504606846976
int		18446744073709551615
*/
