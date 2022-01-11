package main

import (
	"fmt"
	"runtime"
)

func main() {
	//imprime el tipo de sistema operativo y la arquitectura
	fmt.Println("El tipo de sistema Operativo:", runtime.GOOS)
	fmt.Println("La arquitectura:", runtime.GOARCH)
}
