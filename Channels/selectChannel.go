//select - extrae el valor del canales que este listo

package main

import "fmt"

func main() {
	//canal que manipula valores enteros
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	//enviar
	go enviar(par, impar, salir)

	//recibir
	recibir(par, impar, salir)
	fmt.Println("Finalizando")
}

//enviar
func enviar(p, i, s chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}
	//	close(p)
	//	close(i)
	s <- 0
}

//funcion recibir
func recibir(p, i, s <-chan int) {
	for {
		select {
		//revisa si hay un valor en el canal par y lo imprime
		case v := <-p:
			fmt.Println("desde el canal par: ", v)
		//revisa si hay un valor en el canal par y lo imprime
		case v := <-i:
			fmt.Println("desde el canal impar: ", v)
		case v := <-s:
			fmt.Println("desde el canal salir:", v)
			return
		}
	}
}
