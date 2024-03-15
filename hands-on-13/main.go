package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	switch {
	case x <= 100:
		fmt.Printf("The value is between 0 and 100 %d", x)
	case x > 100 && x <= 200:
		fmt.Printf("The value is between 101 and 200 %d", x)
	default:
		fmt.Printf("The value is between 201 and 250 %d", x)

	}

}
