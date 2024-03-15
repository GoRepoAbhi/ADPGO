package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This code is another use of if else statement")
}

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	if x < 4 && y < 4 {
		fmt.Println("They are both less than 40%")
	} else if x > 6 && y > 6 {
		fmt.Println("They are both greater than 60%")
	} else if x >= 4 && x <= 6 {
		fmt.Printf("X is %d %", x*10)
	} else if y != 5 {
		fmt.Println("Y is not 50%")
	} else {
		fmt.Println("I am tired of trying to find statements that makes sense")
	}
}
