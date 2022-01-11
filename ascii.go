package main

import "fmt"

func main() {
	//Tabla para imprimir la tabla del código ASCII
	for i := 32; i <= 123; i++ {
		fmt.Printf("%d\t %#x\t %#U\t\n", i, i, i)
	}
}
