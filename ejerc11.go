package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\t", a, a, a)
	//bit shifting --> desplaza el bit más significativo
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\t", b, b, b)
}
