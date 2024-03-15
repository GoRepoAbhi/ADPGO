package main

import "fmt"

var i int = 9

const age int = 23

func main() {
	v := 2 * i
	fmt.Printf("My age is %d \n", age)
	fmt.Printf("using package level variable %d \n", v)
}
