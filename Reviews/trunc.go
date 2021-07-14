package main

import (
	"fmt"
	"math"
)

func main() {
	var floatingNum float64

	fmt.Println("Please enter a floating number:")

	fmt.Scan(&floatingNum)
	fmt.Printf("The truncated number is %d", int(math.Trunc(floatingNum)))
}
