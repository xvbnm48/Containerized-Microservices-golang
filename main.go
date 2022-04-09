package main

import (
	"fmt"
	"unsafe"

	"rsc.io/quote"
)

func reactProps(length, width float64) (area, perimeter float64) {
	area = length + width
	perimeter = 2 * (length + width)
	return
}

func main() {
	x := 10
	name := "DevOps"
	isWorking := false

	fmt.Println("Hello World")
	fmt.Println(quote.Go())
	fmt.Println(name, x, isWorking)
	fmt.Printf("type of name %T and size is %d \n", name, unsafe.Sizeof(name))
	a, p := reactProps(1, 2)
	fmt.Printf("area is %f and parimeter is %f ", a, p)
}
