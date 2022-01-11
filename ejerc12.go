package main

import "fmt"

func main() {
	//string literal no interpretado ` ` respeta el string que se escribe
	a := `esto es un string 
literal
no interpretado
"¿Lo veo?"`

	fmt.Println(a)

}
