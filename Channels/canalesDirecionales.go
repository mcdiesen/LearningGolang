//canales direccionales
//se especifica cuando es send only - solo envia
//pueden ser receive only - para indicar que solo recibe el canal

//ca := make(<-chan int, 2) esto seria un canal de recepción
package main

import "fmt"

func main() {
	//indica que el canal es send only - solo enviará información
	c := make(chan int)       //canal bidirecional
	cr := make(<-chan int)    //canal receptor
	cs := make(chan<- int, 2) //canal transmisor

	//ca <- 42                  //envia info al canal
	//ca <- 43
	//No podria hacer la siguiente operación
	//fmt.Println(<-ca)
	//fmt.Println(<-ca)

	fmt.Println("---------")
	//imprime el tipo de canal
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cs)
	fmt.Printf("%T\n", cr)

	//cambiar para que sea un canal de solo sendonly

	//de un canal general a especifico si pueden hacer asignaciones

	/* las siguientes conversiones daria un error
	fmt.Println("---------")
	fmt.Println("c\t %T\n", (chan int(cr)))
	fmt.Println("c\t %T\n", (chan int(cs)))
	*/
}
