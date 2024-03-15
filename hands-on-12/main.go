package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	if x >= 0 && x <= 100 {
		fmt.Printf("the value is between 0 and 100 %d", x)
	} else if x > 100 && x <= 200 {
		fmt.Printf("the value is between 101 and 200 %d", x)
	} else {
		fmt.Printf("the value is between 201 and 250 %d", x)
	}

}
