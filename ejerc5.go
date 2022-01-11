package main

import "fmt"

type numero string

var x numero

func main() {
	fmt.Println(x)
	fmt.Printf("El tipo de x es %T\n", x)

	x := 42
	fmt.Println(x)
	//conversion de tipo
	var y = int(x)
	fmt.Println(y)
	fmt.Printf("El tipo de y es %T\n", y)
}
