package main

import (
	"flag"
	"fmt"
)

func main() {
	var inputVal int
	flag.IntVar(&inputVal, "Input", 1, "please type in the value to increase 10")
	flag.Parse()

	fmt.Println("the result is : ", inputVal+10)
}
