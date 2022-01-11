//goroutine que recibe de la goroutine anterior
package main

import "fmt"

func main() {
	//declarar un canal con Buffer (solo goroutine en cada extremo que TX y RX)
	//Goroutine principal
	ca := make(chan int, 2)

	ca <- 42
	ca <- 43
	//se envia al canal los valores 42 y 43 y son impresos
	fmt.Println(<-ca)
	fmt.Println(<-ca)
}
