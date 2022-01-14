//Channels - ejercicio práctico #1 - Canales
//si solo hay una goroutine se bloquea el programa
//se deben utilizar patrones de diseño para que el código no se bloquee
package main

import "fmt"

func main() {
	//canal con buffer, se puede ejecutar sin una funcion goroutine
	c := make(chan int, 1)

	go func() {
		//para evitar que el código se bloquee se debe ejecutar dentro de la funcion
		c <- 42
	}()
	fmt.Println(<-c)

}
