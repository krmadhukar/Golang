package main

import (
	"fmt"
	"reflect"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println(quote.Go())

	// float
	var3 := 1.55

	// boolean
	var4 := true

	// shorthand string array declaration
	var5 := []string{"foo", "bar", "baz"}

	// map is reference datatype
	var6 := map[int]string{100: "Ana", 101: "Lisa", 102: "Rob"}

	// complex64 and complex128
	// is basic datatype
	var7 := complex(9, 15)

	fmt.Printf("var3 = %T\n", var3)
	fmt.Printf("var4 = %T\n", var4)
	fmt.Printf("var5 = %T\n", var5)
	fmt.Printf("var6 = %T\n", var6)
	fmt.Printf("var7 = %T\n", var7)

	// using TypeOf() method of reflect package
	// to determine the datatype of the variables
	fmt.Println()
	fmt.Println("Using reflect.TypeOf Function")
	fmt.Println()

	fmt.Println("var3 = ", reflect.TypeOf(var3))
	fmt.Println("var4 = ", reflect.TypeOf(var4))
	fmt.Println("var5 = ", reflect.TypeOf(var5))
	fmt.Println("var6 = ", reflect.TypeOf(var6))
	fmt.Println("var7 = ", reflect.TypeOf(var7))

	fmt.Println("var3 = ", var3)
}
