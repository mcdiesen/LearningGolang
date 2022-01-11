//usando canales en funciones

package main

import (
	"fmt"
)

func main() {
	//canal general
	c := make(chan int)

	//goroutine - funcion enviar
	go enviar(c)

	//funcion recivir
	recibir(c)

	fmt.Println("Finalizando")
}

//funcion de tipo channel de tipo recibir
func enviar(c chan<- int) {
	c <- 42
}

//funcion canal de tipo receive only
func recibir(c <-chan int) {
	fmt.Println(<-c)
}
