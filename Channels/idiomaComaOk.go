//idioma Coma Ok
//podemos escribir codigo eficiente en canales

package main

import "fmt"

func main() {
	//canal que manipula valores enteros
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan bool)

	//enviar
	go enviar(par, impar, salir)

	//recibir
	recibir(par, impar, salir)
	fmt.Println("Finalizando")
}

//enviar
func enviar(par, impar, salir chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			par <- j
		} else {
			impar <- j
		}
	}
	//	close(p)
	//	close(i)
	salir <- 0
}

//funcion recibir channel
func recibir(par, impar <-chan int, salir <-chan bool) {
	for {
		select {
		//revisa si hay un valor en el canal par y lo imprime
		case v := <-par:
			fmt.Println("desde el canal par: ", v)
		//revisa si hay un valor en el canal par y lo imprime
		case v := <-impar:
			fmt.Println("desde el canal impar: ", v)
		case i, ok := <-salir:
			if !ok {
				fmt.Println("desde coma ok", i, ok)
				return
			} else {
				fmt.Println("desde coma ok:", i, ok)
				return
			}
		}
	}
}
