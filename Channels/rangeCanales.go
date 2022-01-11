//range para iterar sobre canales
//close para cerrar el canal y evitar el bloqueo del código

package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		//envia valores al canal
		for i := 0; i < 5; i++ {
			c <- i
		}
		//cierra el canal
		close(c)
	}()

	//imprimir los valores del canal
	for v := range c {
		fmt.Println(v)
	}
}
