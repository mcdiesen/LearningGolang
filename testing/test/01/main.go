//Package test main
package main

func main() {
	//fmt.Println("La suma de 2+4 es: ", miSuma(2, 4))
	//fmt.Println("La suma de 6+4 es: ", miSuma(6, 4))

	//fmt.Println("La suma de 10+4 es: ", miSuma(10, 4))
	miSuma(2, 3, 4)
}

func miSuma(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
