// Printing and Loggin de mensajes de error
//logging de un error (fmt.Prinln(), log.Println(), log.Fatalln(), log.)
package main

import (
	"fmt"
	"os"
)

func main() {

	err, os:=os.Open("sin-archivo.txt")
	if err != nil{
		fmt.Println("")
	}
}
