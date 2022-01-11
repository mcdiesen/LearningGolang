//Canales - Tipos de Canales (con buffer y sin buffer)
package main

import "fmt"

func main() {
	//declarar un canal sin Buffer (solo goroutine en cada extremo que TX y RX)
	//Goroutine principal
	ca := make(chan int)

	ca <- 42
	//si solo hay un canal se bloquea, requiere goroutine para que funcione
	fmt.Println(<-ca)
}
